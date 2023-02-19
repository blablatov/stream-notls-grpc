###Создание Docker контейнера для gRPC-сервера (build container of server):      

```shell script
docker build -t service .
```

Развернуть задание с серверным gRPC-приложением:         

```shell script
kubectl apply -f grpc-service.yaml
```
