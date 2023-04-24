.PHONY:
.SILENT:

export API_V1 = api/openapi/competency-constructor-v1.yaml

gen-openapi-v1:
	oapi-codegen -generate types -o internal/core/adapter/driver/rest/v1/openapi_type.gen.go -package v1 $$API_V1
	oapi-codegen -generate chi-server -o internal/core/adapter/driver/rest/v1/openapi_server.gen.go -package v1 $$API_V1

gen-api-v1:
	make gen-openapi-v1
	cp $$API_V1 swagger/v1/competency-constructor-v1.yaml


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

# mongodb
export TEST_DB_URI=mongodb://localhost:27019
export TEST_DB_NAME=test
export TEST_DB_CONTAINER_NAME=test-db

run-test-mongo-db:
	docker run --rm -d -p 27019:27017 --name $$TEST_DB_CONTAINER_NAME -e MONGODB_DATABASE=$$TEST_DB_NAME mongo:latest

stop-test-mongo-db:
	docker stop $$TEST_DB_CONTAINER_NAME

#postgresql
export TEST_PSQL_DB_NAME=test-db-name
export TEST_PSQL_DB_CONTAINER_NAME=test-db
export TEST_PSQL_DB_USER=test-user
export TEST_PSQL_DB_PASSWORD=test-password
export TEST_PSQL_DB_HOST=localhost
export TEST_PSQL_DB_PORT=5432

run-test-psql-db:
	docker run --rm -d -p 5432:5432\
		 --name $$TEST_PSQL_DB_CONTAINER_NAME\
		 -e POSTGRES_DB=$$TEST_PSQL_DB_NAME -e POSTGRES_USER=$$TEST_PSQL_DB_USER\
 		 -e POSTGRES_PASSWORD=$$TEST_PSQL_DB_PASSWORD postgres:latest

stop-test-psql-db:
	docker stop $$TEST_DB_CONTAINER_NAME


#migrate
create-migration:
	 migrate create -ext sql -dir migrations ${NAME}

dev-up:
	migrate -path ./migrations -database\
		'postgres://${TEST_PSQL_DB_USER}:${TEST_PSQL_DB_PASSWORD}@${TEST_PSQL_DB_HOST}:${TEST_PSQL_PORT}/${TEST_PSQL_DB_NAME}?sslmode=disable' up

dev-down:
	migrate -path ./migrations -database\
		'postgres://${TEST_PSQL_DB_USER}:${TEST_PSQL_DB_PASSWORD}@${TEST_PSQL_DB_HOST}:${TEST_PSQL_PORT}/${TEST_PSQL_DB_NAME}?sslmode=disable' down 1