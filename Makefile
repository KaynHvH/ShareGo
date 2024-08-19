build:
	@go build -o bin/ShareGo

run: build
	@./bin/ShareGo

test:
	@go test -v ./...