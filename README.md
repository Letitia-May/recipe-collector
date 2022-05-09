# Recipe collector

A simple app to help me keep all the recipes I've cooked and loved in one place, while also learning Go! :woman_cook: :woman_technologist:

## Getting Started

### Dependencies

* Go version `go1.17.6` and above

### Installing

Run the following command from the root directory to get all the required dependencies:
```
go get .
```
Set the environment variables for accessing the database:
```
export DBUSER=<username>
export DBPASS=<password>
export DBADDR=<address>
```

### Running

Note: Currently just a command line app.

From the root directory to get a list of recipes:
```
go run .
```
Access `http://localhost:8080/recipes` to see list of all recipes.