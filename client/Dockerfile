# Multi-stage client build
# Многоэтапная сборка client

FROM golang AS build

ENV location /go/src/github.com/blablatov/stream-notls-grpc

WORKDIR ${location}/client

ADD ./main.go ${location}/client

RUN go mod init github.com/blablatov/stream-notls-grpc/client

RUN CGO_ENABLED=0 go build -o client

# Go binaries are self-contained executables. Используя директиву FROM scratch - 
# Go образы  не должны содержать ничего, кроме одного двоичного исполняемого файла.
FROM scratch
COPY --from=build ./client ./client

ENTRYPOINT ["./client"]
EXPOSE 50051