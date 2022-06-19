app-test:
	go test tests

app-build:
	go build -o build/api/main cmd/api

app-run:
	go run ./cmd/api/main.go
