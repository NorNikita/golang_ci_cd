.PHONY: lint
lint:
	golangci-lint run -v ./app/...

.PHONY: build
build:
	go build -o ./bin/app ./code

.PHONY: test
test:
	go test -v -race ./code