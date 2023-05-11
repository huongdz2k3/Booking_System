FROM golang:1.20.3-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/main.go
# Set the default command to start the server
CMD ["go", "run", "./cmd" ,"api"]
