# Multi-stage server build
# Многоэтапная сборка service

FROM golang AS build

ENV location /go/src/github.com/blablatov/stream-notls-grpc

WORKDIR ${location}/service

ADD ./main.go ${location}/service
ADD ./proto  ${location}/proto

#ADD ./go.mod  ${location}/service
#ADD ./go.sum  ${location}/service
RUN go mod init github.com/blablatov/stream-notls-grpc/service

# Download all the dependencies. Загрузка зависимостей
RUN go get -d ./...
# Install the package. Установка пакетов
RUN go install ./...

# Build server app. Сборка серверного приложения.
RUN CGO_ENABLED=0 go build -o /bin/service 

# Go binaries are self-contained executables. Используя директиву FROM scratch - 
# Go образы  не должны содержать ничего, кроме одного двоичного исполняемого файла.
FROM scratch
COPY --from=build /bin/tls-service /bin/service

ENTRYPOINT ["/bin/service"]
EXPOSE 50051