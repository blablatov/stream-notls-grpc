[![Go](https://github.com/blablatov/stream-notls-grpc/actions/workflows/stream-notls-grpc.yml/badge.svg)](https://github.com/blablatov/stream-notls-grpc/actions/workflows/stream-notls-grpc.yml)
### Building and Running gRPC service

In order to build, Go to ``Go`` module directory location `stream-notls-grpc/service` and execute the following
 shell command:
```
go build -v
./service
```  

### Building and Running gRPC client   

In order to build, Go to ``Go`` module directory location `stream-notls-grpc/client` and execute the following shell command:
```
go build -v 
./client
```  

### Generates Server and Client side code via proto-file     
Go to ``Go`` module directory location `stream-notls-grpc/proto` and execute the following shell commands:    
``` 
protoc order_management.proto --go_out=./ --go-grpc_out=./
protoc order_management.proto --go-grpc_out=require_unimplemented_servers=false:.
``` 
