# Запуск

В одном терминале
```shell
git clone https://github.com/syalsr/axitex.git
cd axitex/cmd
go run server.go 50051 mul
```

В другом терминале
```shell
cd axitex/cmd
go run client.go 127.0.0.1:50051 25 35
```