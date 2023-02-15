module github.com/blablatov/stream-notls-grpc

go 1.18

replace github.com/blablatov/stream-notls-grpc/proto => ./proto

replace github.com/blablatov/stream-notls-grpc/mockups => ./mockups

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	cloud.google.com/go/compute v1.15.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/golang/mock v1.6.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/oauth2 v0.4.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
)
