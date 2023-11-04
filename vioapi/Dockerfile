# Build stage
FROM golang:1.20-alpine AS builder
ARG CUJOO_TOKEN 
RUN apk add --update ca-certificates build-base git
RUN git config --global url."https://cujoo:$CUJOO_TOKEN@github.com/".insteadOf "https://github.com/"
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o application-internal -tags musl ./cmd/internal/main.go
RUN GOOS=linux GOARCH=amd64 go build -o application-external -tags musl ./cmd/external/main.go

# Final stage
FROM alpine:3.16
RUN apk add --update ca-certificates
COPY --from=builder /build/application-internal /usr/bin/application-internal
COPY --from=builder /build/application-external /usr/bin/application-external