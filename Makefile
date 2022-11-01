.PHONY:
.SILENT:

build:
	go mod download && CGO_ENABLED=0 go build -o ./.bin/competency-constructor ./cmd/competency-constructor

lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

run:
	go run ./cmd/competency-constructor


export TEST_DB_NAME=test
export TEST_DB_CONTAINER_NAME=test-db
export TEST_DB_USER=test-user
export TEST_DB_PASSWORD=test-password

run-test-db:
	docker run --rm -d -p 5432:5432 --name $$TEST_DB_CONTAINER_NAME -e POSTGRES_DB=$$TEST_DB_CONTAINER_NAME  -e POSTGRES_USER=$$TEST_DB_USER  -e POSTGRES_PASSWORD=$$TEST_DB_PASSWORD postgres:latest

stop-test-db:
	docker stop $$TEST_DB_CONTAINER_NAME