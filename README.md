# fx-bank
Backend server for fx-bank

- To set the project up, an `.env` file has to be created in the root directory. `.env_test` can be used as a clone for what `.env` must contain.
- Run `make migrate-up` in root directory to run migrations against the database.
- Run server with `go run cmd/main.go` in root directory.
