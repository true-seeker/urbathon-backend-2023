## Бэкенд приложения "Городские связи" для хакатона "Урбатон-2023"

Используемые технологии: Golang, PostgreSQL, Redis, s3 (Yandex Object Storage).
Пакеты Go: 
- [Gin](https://github.com/gin-gonic/gin) (веб-фреймворк)
- [go-jet](https://github.com/go-jet/jet) (кодогенерация моделей БД, и конструктор запросов БД)
- [go-migrate](https://github.com/golang-migrate/migrate) (миграции БД)
- aws (s3)
- [swag](https://github.com/swaggo/swag) (генерация спецификации Swagger)

#### Локальный запуск:

Необходимо, чтобы на хосте были свободны порты `8080`,`5633`, `6379`
После запуска поднимется инстанс PostgreSQL, Redis и приложения с бэкендом

```shell
cd deployments
docker-compose up -d 
```

После запуска спецификация Swagger будет доступа по ссылке: http://localhost:8080/swagger/index.html#/

### Авторизационные данные
После локального запуска создается аккаунт администратора системы
Email: admin@urbathon.ru 
Пароль: admin
