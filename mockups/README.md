### Тестирование функциональность клиентского кода без подключения к серверу. 
### Testing code without conn to server            

### Go 1.16+
```shell script
go install github.com/golang/mock/mockgen@v1.5.0
```

### Использование Gomock, генерация макетов интерфейсов клиентского gRPC-приложения. Use Gomock      
Для генерации макета интерфейса OrderManagementClient, выполнить.   
(Runs generation code of mock up for interface OrderManagementClient): 
       
```shell script
mockgen github.com/blablatov/stream-notls-grpc/proto OrderManagementClient > mockups/order_mock.go
```

### Run test    

```shell script
./mockups/order_mock_test.go
```


