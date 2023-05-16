#Build stage
FROM golang:1.20.3-alpine
# cmd below is to create a directory in the container
WORKDIR /app
# cmd below is to copy the go.mod file from the current directory to the container
COPY go.mod .
COPY go.sum .
# cmd below is to run the go mod tidy command
RUN go mod tidy
# cmd below is to copy all the files from the current directory to the container
COPY . .
# cmd below is to build the go file in the container and create an executable file
# -o server is to create an executable file named server
# cmd/main.go is to specify the file to be built
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/main.go


# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
# cmd below is to specify the executable file to be run when the container is started
CMD ["./server", "api"]
