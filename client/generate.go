package client

//go:generate protoc -I=proto --go_out=:pb --go-grpc_out=:pb proto/*.proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative
