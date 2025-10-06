package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/go-sql-driver/mysql"
)

func runMigrations() error {
	// Capture connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("DBADDR"),
		DBName: "recipe_collector",
	}

	// Get a db handle
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("✅ Connected to database successfully")

	// Create migrations table if it doesn't exist
	if err := createMigrationsTable(db); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Run migrations
	if err := executeMigrations(db); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	fmt.Println("🎉 All migrations completed successfully!")
	return nil
}

func createMigrationsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`
	_, err := db.Exec(query)
	return err
}

func executeMigrations(db *sql.DB) error {
	// Get list of applied migrations
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	// Get list of migration files
	migrationFiles, err := getMigrationFiles()
	if err != nil {
		return fmt.Errorf("failed to get migration files: %w", err)
	}

	// Run pending migrations
	for _, file := range migrationFiles {
		version := strings.TrimSuffix(file, ".sql")
		
		if _, applied := appliedMigrations[version]; applied {
			fmt.Printf("⏭️ Skipping already applied migration: %s\n", file)
			continue
		}

		fmt.Printf("🔄 Running migration: %s\n", file)
		
		if err := runMigrationFile(db, file); err != nil {
			return fmt.Errorf("failed to run migration %s: %w", file, err)
		}

		if err := markMigrationAsApplied(db, version); err != nil {
			return fmt.Errorf("failed to mark migration as applied %s: %w", version, err)
		}

		fmt.Printf("✅ Successfully applied migration: %s\n", file)
	}

	return nil
}

func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := make(map[string]bool)
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		applied[version] = true
	}

	return applied, rows.Err()
}

func getMigrationFiles() ([]string, error) {
	var files []string
	
	err := filepath.WalkDir("migrations", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".sql") && !strings.HasSuffix(d.Name(), "_down.sql") {
			files = append(files, d.Name())
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}

	// Sort files to ensure consistent order
	sort.Strings(files)
	return files, nil
}

func runMigrationFile(db *sql.DB, filename string) error {
	// Read migration file
	content, err := os.ReadFile(filepath.Join("migrations", filename))
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Parse SQL statements
	statements := parseSQLStatements(string(content))
	
	// Execute each statement
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue // Skip empty statements
		}

		if _, err := db.Exec(stmt); err != nil {
			return fmt.Errorf("failed to execute statement '%s': %w", stmt, err)
		}
	}

	return nil
}

func parseSQLStatements(content string) []string {
	var statements []string
	var currentStatement strings.Builder
	
	lines := strings.Split(content, "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "--") {
			continue
		}
		
		// Add line to current statement
		if currentStatement.Len() > 0 {
			currentStatement.WriteString(" ")
		}
		currentStatement.WriteString(line)
		
		// If line ends with semicolon the statement is complete
		if strings.HasSuffix(line, ";") {
			stmt := strings.TrimSuffix(currentStatement.String(), ";")
			if stmt != "" {
				statements = append(statements, stmt)
			}
			currentStatement.Reset()
		}
	}
	
	// Handle case where last statement doesn't end with semicolon
	if currentStatement.Len() > 0 {
		stmt := currentStatement.String()
		if stmt != "" {
			statements = append(statements, stmt)
		}
	}
	
	return statements
}

func markMigrationAsApplied(db *sql.DB, version string) error {
	_, err := db.Exec("INSERT INTO schema_migrations (version) VALUES (?)", version)
	return err
}