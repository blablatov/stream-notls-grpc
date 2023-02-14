## Service and Client - Go Implementation

## Building and Running Service

In order to build, Go to ``Go`` module directory location `stream-grpc/service` or `stream-grpc/tls-service` and execute the following
 shell command:
```
go build -i -v -o server (tls-server) (mtls-server)
```

In order to run, Go to ``Go`` module directory location `stream-grpc/service` `stream-grpc/service` or `stream-grpc/tls-service` and execute the following
shell command:

```
./server
./tls-server
./mtls-server
```

## Building and Running Client   

In order to build, Go to ``Go`` module directory location `stream-grpc/client` or `stream-grpc/tls-client` and execute the following shell command:
```
go build -i -v -o client (tls-client) (mtls-server)
```

In order to run, Go to ``Go`` module directory location `stream-grpc/client` or `stream-grpc/tls-client` and execute the following shell command:

```
./client
./tls-client
./mtls-client
```

## Additional Information

### Generate Server and Client side code   
Go to ``Go`` module directory location `stream-grpc/proto` and execute the following shell commands:    
``` 
protoc product_info.proto --go_out=./ --go-grpc_out=./
protoc product_info.proto --go-grpc_out=require_unimplemented_servers=false:.
``` 
