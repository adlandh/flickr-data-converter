build:
	@go build -o flickr-data-converter ./cmd
run-linter:
	@golangci-lint run
