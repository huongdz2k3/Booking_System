# Build stage
FROM golang:1.20.3-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
CMD ["./server", "api"]


