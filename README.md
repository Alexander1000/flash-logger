# flash-logger

Сборка
```shell
make
```

Запуск
```shell
./bin/flash-logger -c etc/config.json
```

Отправка TCP пакета
```shell
curl -X POST -H 'Authorization: Bearer fjldf7sfSf8asLf8ss7' -d '{"limit":20,"offset":0}' http://localhost:42234/1/logs
```

Отправка пакета по UDP
```shell
echo -n "hello world" | nc -4u -w1 localhost 42234
```
