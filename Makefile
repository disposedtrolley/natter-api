ENTRYPOINT:=github.com/disposedtrolley/natter-api/cmd/natter-api
BINARY:=natter-api

build:
	go build -o $(BINARY) $(ENTRYPOINT)

run:
	go run $(ENTRYPOINT)

test:
	go test -race -cover ./...
