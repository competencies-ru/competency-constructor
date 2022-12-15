.PHONY:
.SILENT:

export API_V1 = api/openapi/competency-constructor-v1.yaml

gen-api-v1:
	oapi-codegen -generate types -o internal/core/adapter/driver/rest/v1/openapi_type.gen.go -package v1 $$API_V1
	oapi-codegen -generate chi-server -o internal/core/adapter/driver/rest/v1/openapi_server.gen.go -package v1 $$API_V1

build:
	go mod download && CGO_ENABLED=0 go build -o ./.bin/competency-constructor ./cmd/competency-constructor

lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .

run:
	go run ./cmd/competency-constructor


export TEST_DB_NAME=test-db-name
export TEST_DB_CONTAINER_NAME=test-db
export TEST_DB_USER=test-user
export TEST_DB_PASSWORD=test-password
export TEST_DB_HOST=localhost

run-test-db:
	docker run --rm -d -p 5432:5432\
		 --name $$TEST_DB_CONTAINER_NAME\
		 -e POSTGRES_DB=$$TEST_DB_NAME -e POSTGRES_USER=$$TEST_DB_USER\
 		 -e POSTGRES_PASSWORD=$$TEST_DB_PASSWORD postgres:latest

stop-test-db:
	docker stop $$TEST_DB_CONTAINER_NAME


#migrate

create-migration:
	goose -dir=migrations create ${FILE_NAME} ${TYPE}

dev-up:
	goose -dir ./migrations postgres\
 	"host=${TEST_DB_HOST} user=${TEST_DB_USER} password=${TEST_DB_PASSWORD} dbname=${TEST_DB_NAME} sslmode=disable" up
