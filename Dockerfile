FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

COPY authentication/go.mod ./authentication
COPY authentication/go.sum ./authentication
COPY shared ./shared

RUN go mod download

# Copy the code into the container
COPY authentication ./authentication

# Build the application
RUN go build -o authentication/cmd/main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 6666

# Command to run when starting the container
CMD ["/dist/main"]