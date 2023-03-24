migrate-up:
	cd internal/migrations && goose postgres "postgresql://postgres:postgres@localhost:5432/fx?sslmode=disable" up
.PHONY: migrate-up

migrate-down:
	cd internal/migrations && goose postgres "postgresql://postgres:postgres@localhost:5432/fx?sslmode=disable" down
.PHONY: migrate-down