FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    Environment=docker \
    BILL_HTTP_PORT=0.0.0.0:7777 \
    BILL_GRPC_PORT=0.0.0.0:8888 \
    AMQP_URL=amqp://guest:guest@host.docker.internal:5672/ \
    JWTKey=TopSecretKey \
    RefreshKey=RefreshTopSecretKey \
    Redis_Host=redis:6379 \
    DBHost=host.docker.internal \
    BILLS_DBDatabase=naerabills \
    DBUser=root \
    DBPassword=password \
    DBPort=3306 \
    RedisPass=password \
    TwilloSID=AC56464c412d3a5e05eee77f45c741d912 \
    TwilloToken=a5a84f35e80adcde4987cf087a67d89b \
    TwilloPhone=+14436489834 \
    FL_SECRETKEY_LIVE=FLWSECK-2795fda29f319fba2067279cb8300ecb-X \
    FL_KEY_LIVE=2795fda29f314b1e28585d75 \
    FL_SECRETKEY_TEST=FLWSECK_TEST-be6475503d295c1be0b10ee8e971671f-X \
    FL_KEY_TEST=FLWSECK_TESTc5fc82d7d2b7 

# Move to working directory /build
WORKDIR /build

COPY bills/go.mod /build/bills/
COPY bills/go.sum /build/bills/

COPY shared/go.mod /build/shared/
COPY shared/go.sum /build/shared/

WORKDIR /build/shared
RUN go mod tidy

WORKDIR /build/bills
RUN go mod tidy

WORKDIR /build

# Copy the code into the container
COPY shared /build/shared/
COPY bills /build/bills/

# Build the application

WORKDIR /build/bills
RUN go build -o naera_bills cmd/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/bills/naera_bills .

# Export necessary port
EXPOSE 5555
EXPOSE 7777
EXPOSE 6379

# Command to run when starting the container
CMD ["/bin/sh", "-c", "/dist/naera_bills"]