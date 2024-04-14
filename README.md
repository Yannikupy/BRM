# BRM

## Документация

* [Техническое задание](https://docs.google.com/document/d/1NK2MnQOKXtqQmShMxQeJxH1kPGQj8I0yIqQ5TaXJ1d0/edit#heading=h.bwhogp1n8aje)
* [Описание сервисов](https://docs.google.com/document/d/1NK2MnQOKXtqQmShMxQeJxH1kPGQj8I0yIqQ5TaXJ1d0/edit#heading=h.xn76w1oe0rlr)

## Запуск сервисов

Чтобы поднять все сервисы и их зависимости (для разработки фронта):

```shell
make run_dev_front
```

Чтобы поднять только зависимости сервисов (для разработки самих сервисов)

```shell
make run_dev_back
```

## swagger-документация

* [transport-api](http://localhost:8090/swagger/index.html)
* [registration](http://localhost:8091/swagger/index.html)
* [auth](http://localhost:8092/swagger/index.html)
* [images](http://localhost:8093/swagger/index.html)
