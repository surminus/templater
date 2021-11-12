.PHONY=build
build:
	@go build -o bin/templater

.PHONY=test
test:
	@go test
