# flash-logger

Сборка
```shell
make
```

Запуск
```shell
./bin/flash-logger -c etc/config.json
```

Отправка пакета по UDP
```shell
echo -n "hello world" | nc -4u -w1 localhost 42234
```
