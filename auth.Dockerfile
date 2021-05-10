FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    Environment=production \
    HTTP_PORT=0.0.0.0:5555 \
    GRPC_PORT=0.0.0.0:6666 \
    AMQP_URL=amqp://guest:guest@host.docker.internal:5672/ \
    JWTKey=TopSecretKey \
    RefreshKey=RefreshTopSecretKey \
    Redis_Host=host.docker.internal:6379 \
    DBHost=host.docker.internal \
    DBDatabase=naeraauth \
    DBUser=root \
    DBPassword=password \
    DBPort=3306 \
    RedisPass=password \
    TwilloSID=AC56464c412d3a5e05eee77f45c741d912 \
    TwilloToken=a5a84f35e80adcde4987cf087a67d89b \
    TwilloPhone=+14436489834

# Move to working directory /build
WORKDIR /build

COPY authentication/go.mod /build/authentication/
COPY authentication/go.sum /build/authentication/


COPY shared/go.mod /build/shared/
COPY shared/go.sum /build/shared/

WORKDIR /build/shared
RUN go mod tidy

WORKDIR /build/authentication
RUN go mod tidy


WORKDIR /build

# Copy the code into the container
COPY shared /build/shared/
COPY authentication /build/authentication/

# Build the application
WORKDIR /build/authentication
RUN go build -o naera_auth cmd/main.go


# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/authentication/naera_auth .

# Export necessary port
EXPOSE 5555
EXPOSE 6379
EXPOSE 5672

# Command to run when starting the container
CMD ["/bin/sh", "-c", "/dist/naera_auth"]