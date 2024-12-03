# gRPC Client and Server using Golang Workspace

## create folder structure

```bash
mkdir grpcchat
cd grpcchat
mkdir chat
go mod init grpcchat
cd ..
mkdir client
go mod init client
cd ..
mkdir server
go mod init server
cd ..
go work init
go work use ./grpcchat
go work use ./client
go work use ./server
```

## Create gRPC proto file and generate Go code

  create `grpcchat/chat.proto` file

```bash
cd grpcchat
protoc --go_out=. --go-grpc_out=. chat.proto
```

## Create gRPC server

- create `server/chatimpl/chat.go` file
  - implement `chat.Server` interface
    - dont forget to embed `chat.UnimplementedChatServiceServer`
  - implement `chat.ChatServiceServer` interface
- create server.go file
  - implement `main` function

## Create gRPC client

- create `client/client.go` file
  - implement `main` function
