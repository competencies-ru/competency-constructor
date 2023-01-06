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


linux-build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/competency-constructor ./cmd/competency-constructor


run-dev: linux-build
	docker-compose -f ./deployments/dev/docker-compose.yml --project-directory . up --remove-orphans competency-constructor

stop-dev:
	docker-compose -f ./deployments/dev/docker-compose.yml --project-directory . stop competency-constructor

export TEST_DB_URI=mongodb://localhost:27019
export TEST_DB_NAME=test
export TEST_DB_CONTAINER_NAME=test-db

run-test-db:
	docker run --rm -d -p 27019:27017 --name $$TEST_DB_CONTAINER_NAME -e MONGODB_DATABASE=$$TEST_DB_NAME mongo:latest

stop-test-db:
	docker stop $$TEST_DB_CONTAINER_NAME

