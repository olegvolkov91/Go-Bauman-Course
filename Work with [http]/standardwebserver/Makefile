.SILENT:
.PHONY:

build:
	go build -v ./cmd/api

run: build
	./api

migrate:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" up

drop:
	migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" down
