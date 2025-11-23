# Go.Back - Бэкенд для My Travel Map

Приложение написанное на go для работы с полигонами и картами. Путешествуй закаршивая новые места и отмечайся где уже побывал!

## 📁 Запуск проекта

 ### Режим разработки:

 - Установить зависимости:

 ```bash
 $ go mod download
 ```

 - В проекте есть миграции, поэтому перед запуском проекта нужно подготовить свои переменные окружения из example.env создав .env файл, запустить PostgreSQL и накатить туда миграции (я использую migrate):

 ```bash
$ migrate -path ./migrations/postgres -database 'postgres://lastmile:<пароль из .env>@localhost:5433/go?sslmode=disable' up
```

 - Запустить проект:

```bash
$ go run cmd/main.go
```
