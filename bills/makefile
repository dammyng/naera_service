-include cmd/.env
export 


build:
	go build -o cmd/bin/naera_bills cmd/main.go

run:
	cd cmd && go run main.go

test:
	go test ./... -v

compile:
	echo "Compiling for linux"
	env GOOS=linux go build -o cmd/bin/naera_bills cmd/main.go

proto:
	protoc -I=./ --proto_path=models/party --go_out=plugins=grpc:./ models/v1/billing.proto
