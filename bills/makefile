-include cmd/.env
export 


build:
	go build -o cmd/bin/naera_auth cmd/main.go

run:
	go run cmd/main.go

test:
	go test ./...

compile:
	echo "Compiling for linux"
	env GOOS=linux go build -o cmd/bin/naera_bills cmd/main.go

proto:
	protoc -I=./ --proto_path=models/party --go_out=plugins=grpc:./ models/v1/biller.proto 
