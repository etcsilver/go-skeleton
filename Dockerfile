FROM golang:1.18.1-alpine as builder

# Move to working directory /build
WORKDIR /build

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the application
RUN go build -buildvcs=false -o main ./cmd/

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .
RUN cp /build/.env .

# Build a small image
#FROM scratch
FROM alpine:3.11

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

RUN apk add --no-cache tzdata    

COPY --from=builder /dist/main /
COPY --from=builder /dist/.env /

# Command to run
ENTRYPOINT ["/main"]