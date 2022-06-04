build:
	go build -o bin/ibanator cmd/ibanator/main.go

run:
	go run cmd/ibanator/main.go

test:
	go test -v ./...
