GOPACKAGES = $(shell go list ./...)
MIGRATIONS_DIR=$(shell pwd)/db/migrations
DB_DSN=postgres://postgres:postgres@host.docker.internal:5432/vio?sslmode=disable

create-test-db:
	-docker stop postgres
	-docker rm postgres
	docker run --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=vio -p 5432:5432 -d postgis/postgis

migrate-up:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_DSN) up

migrate-down:
	docker run --rm -v $(MIGRATIONS_DIR):/migrations --network host migrate/migrate -path=/migrations/ -database $(DB_DSN) down -all

gen-sql:
	sqlc generate

mockgen:
	@echo "Generating mocks..."
	@mockgen -source=storage/storage.go -destination=storage/mock.go -package=storage
	@mockgen -source=db/gen/querier.go -destination=db/gen/mock.go -package=gen
	@mockgen -source=viodata.go -destination=mock.go -package=viodata

test:
	@go test ./...
