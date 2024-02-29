# BRM

## Перед тем как запускать

Для запуска следующих сервисов нужны расширенные конфиги, которых нет в 
репозитории:

* [auth](back/auth/config)
* [brm-ads](back/brm-ads/config)
* [brm-core](back/brm-core/config)
* [transport-api](back/transport-api/config)

В эти директории нужно добавить `.env` файлы, которые соответствуют схеме, 
описанной в `.env.dist`.

## Запуск сервисов

Чтобы поднять все сервисы и их зависимости (для разработки фронта):

```shell
make run_dev_front
```

Чтобы поднять только зависимости сервисов (для разработки самих сервисов)

```shell
make run_dev_back
```
