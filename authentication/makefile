-include cmd/.env
export 

clear:
	rm -r cmd/*

build:
	go build -o cmd/bin/naera-auth cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./...

compile:
	echo "Compiling for linux"
	env GOOS=linux go build -o cmd/bin/naera-auth cmd/main.go