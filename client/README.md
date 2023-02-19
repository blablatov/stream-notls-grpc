### Тестирование функциональность клиентского кода с подключением к серверу. 
### Testing code with conn to server          
  
Традиционный тест, который запускает клиент для проверки удаленного метода сервиса.  
Перед его выполнением запустить grpc-сервер.   
(Conventional test that starts a gRPC client test the service with RPC.Before his execute run grpc-server):      

```shell script
./service/service
```

Для тестирования клиента, без подключения к серверу, выполнить сгенерированный тестовый код.      
(Runs generation code of mock up for interface OrderManagementClient):   
       
```shell script
./mockups/order_mock_test.go
```

Создание Docker контейнера для gRPC-клиента (build container of client):    

```shell script
docker build -t client .
```

Развернуть задание с клиентским gRPC-приложением:    

```shell script
kubectl apply -f grpc-client.yaml
```


