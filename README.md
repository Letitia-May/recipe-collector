# Recipe collector

A simple app to help me keep all the recipes I've cooked and loved in one place, while also learning Go! :woman_cook: :woman_technologist:

## Dependencies

* Go version `go1.17.6` and above

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

Set the environment variables for accessing the database:
```
export DBUSER=<username>
export DBPASS=<password>
export DBADDR=<address>
```

From the `backend` directory start the web server:
```
go run .
```
Visit `http://localhost:8080/recipes` to see the list of all recipes as a json response.

From the `frontend` directory start the local development server:
```
npm run develop
```
Gatsby will start a hot-reloading development environment at `http://localhost:8000`.

## Technology
### Backend
- GoLang
- MySQL
- Chi (router for building Go HTTP services)
### Frontend
- Typescript
- React
- Gatsby (framework for building websites)
- Styled-components