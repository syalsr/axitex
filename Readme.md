# Запуск

В одном терминале
```shell
git clone https://github.com/syalsr/axitex.git
cd axitex/build
./server 50051 mul
```

В другом терминале
```shell
cd axitex/cmd
./client 127.0.0.1:50051 25 35
```