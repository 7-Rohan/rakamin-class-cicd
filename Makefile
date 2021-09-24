install:
	@go mod init hello-makefile

test:
	@go test -v .

run:
	@go run main.go

build:
	@go build