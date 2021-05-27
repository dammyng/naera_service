-include cmd/.env
export 


build:
	go build -o cmd/bin/naera_auth cmd/main.go

run:
	cd cmd && go run main.go

test:
	go test ./...

compile:
	echo "Compiling for linux"
	env GOOS=linux go build -o cmd/bin/naera_auth cmd/main.go

proto:
	protoc -I=./ --proto_path=models/party --go_out=plugins=grpc:./ models/v1/account.proto 
