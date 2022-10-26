.PHONY:
.SILENT:

build:
	go mod download && CGO_ENABLED=0 go build -o ./.bin/competency-constructor ./cmd/competency-constructor

lint:
	golangci-lint run

gofumpt:
	gofumpt -l -w .
