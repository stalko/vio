# VIO - API

This application is responsible for providing access to `viodata` over HTTP. Also it is an entrypoint for importing `data_dump.csv` to a database over the same `viodata` library.

## HTTP server

The entry point is located in `./cmd/vioapi/main.go`. The server implements HTTP endpoint `/ip_location/:ip` where, given IP address, returns information about the IP address' location (e.g. country, city). There is one more endpoint `/swagger/*any` for sharing automatically generated swagger documentation( [How to re-generate swagger documentation](#swagger) ). The server will automatically redirect any GET requests from `/` to `/swagger/index.html` for reading documentation and testing API endpoint.

## Importer

The entry point is located in `./cmd/importer/main.go`. Given CSV file imports to the DB storage valid rows and retrieves statistics about the time elapsed, as well as the number of entries accepted/discarded.

# Testing

The main packages like: `config`, `logging` and `server` are covered with tests. You can execute tests by running command:
```
make test
```

## VS Code

This project includes VS Code configuration to debug locally `api` and `importer`. Do not forget to provide correct `.env` file and place it in the main(`vio/vioapi`) directory of the application. `.env` file can be create from the example file `.env.example` in the root directory `vio`.

# Dependencies

## VIODATA
This is the main dependency of the application. The link to the library is `github.com/stalko/viodata`, but temporary replaced by `../viodata` in `go.mod` file. There is a TODO task in the root README.md file to improve it in the future. It will be nice to have version control for this library, as well.


# swagger
Documentation of the HTTP Server is automatically generated. You can find comments in the `./pkg/server/http.go` file that are used for documentation generation. You can find more about attributes, parameters, annotations and more here: https://github.com/swaggo/swag.
For re-generating documentation ensure you've installed:
```
go get -u github.com/swaggo/swag/cmd/swag
```
After successful installation execute:
```
make gen
```
or

```
swag i -g ./cmd/vioapi/main.go -o ./docs
```

It will re-generate documentation in the `./docs` file.
