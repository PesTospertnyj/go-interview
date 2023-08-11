build:
	go build -o bin/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)
run:
	go run cmd/main.go