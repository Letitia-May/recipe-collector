# Recipe collector

An app to help me keep all the recipes I've cooked and loved in one place, while also learning some new tech! :woman_cook: :woman_technologist:

## Dependencies

-   Go version `go1.19` and above

## Installing

Run the following command from the `backend` directory to get all the required Go dependencies:

```
go get .
```

Run the following command from the `frontend` directory to get all the required JS dependencies:

```
npm install
```

## Developing

### Backend

#### Environment Setup

Set the environment variables for accessing the database:

```
export DBUSER=<username>
export DBPASS=<password>
export DBADDR=<address>
```

#### CLI Commands

The application includes a CLI with the following commands:

- `serve` - Start the web server
- `migrate` - Run database migrations

#### Database Migrations

Before starting the server, run the database migrations from the `backend` directory to set up the required tables:

```
go run . migrate
```

For more detailed information about migrations, see the [migrations README](backend/migrations/README.md).

#### Running the Server

From the `backend` directory start the web server:

```
go run . serve
```

Visit `http://localhost:8080/recipes` to see the list of all recipes as a json response.

Note: If you want to run the server with live reloading, install https://github.com/cosmtrek/air and run `air` instead. E.g. `air serve`.

### Frontend

From the `frontend` directory start the local development server:

```
npm run dev
```

Next.js will start a hot-reloading development environment at `http://localhost:3000`.

## Technology

### Backend

-   GoLang
-   MySQL
-   Chi (router for building Go HTTP services)
-   urfave/cli (CLI framework)

### Frontend

-   Typescript
-   React
-   Next.js
-   ESLint
-   Prettier
