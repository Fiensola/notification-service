.PHONY: run lint fmt test

run:
	go run cmd/main.go

lint:
	golangci-lint run

fmt:
	go fmt ./...

test:
	go test -v ./...