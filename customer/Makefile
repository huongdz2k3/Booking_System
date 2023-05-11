ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
protoc:
	protoc -I=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb