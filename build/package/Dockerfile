FROM golang:alpine AS builder
MAINTAINER Moises Brenes <moises.brenes@gmail.com>

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    DOCKERIZE_VERSION=v0.6.1

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o server cmd/*.go

# Download dockerize
RUN apk add --no-cache openssl
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz
RUN tar -C . -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/server .
RUN cp /build/dockerize .

# Build a small image
FROM scratch

COPY --from=builder /dist/server /
COPY --from=builder /dist/dockerize /

# Copy static and config files
COPY ./configs/.env /
COPY ./web/templates /web/templates

# Export necessary port
EXPOSE 8000

# Command to run when starting the container
CMD ["/dockerize", "-wait", "tcp://mysql:3306", "-timeout", "30s", "/server", "serve"]