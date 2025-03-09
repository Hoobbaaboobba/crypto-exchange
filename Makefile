build:
	@go build -o bin/exchange main.go

run: build
	@./bin/exchange

test:
	@go test -v ./
