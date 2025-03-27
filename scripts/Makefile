include .env
export $(shell sed 's/=.*//' .env)

PROJECT_ROOT := $(realpath $(dir $(abspath $(lastword $(MAKEFILE_LIST))))/..)

# Запуск серверу
run:
	cd $(PROJECT_ROOT) && go run cmd/main.go

# Ініціалізація модуля Go (тільки один раз)
init:
	@if [ ! -f "$(PROJECT_ROOT)/go.mod" ]; then \
		echo "Initializing Go module..."; \
		cd $(PROJECT_ROOT) && go mod init example.com/survey; \
	else \
		echo "go.mod already exists, skipping initialization."; \
	fi

# Встановлення бібліотек
install:
	go get -u github.com/lib/pq 
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres

# Створення бази даних
createdb:
	PGPASSWORD=$(DB_PASSWORD) psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "CREATE DATABASE $(DB_NAME);"

# Видалення бази даних
dropdb:
	PGPASSWORD=$(DB_PASSWORD) psql -U $(DB_USER) -h $(DB_HOST) -p $(DB_PORT) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

# Міграція (створення таблиць)
migrate:
	go run migrate/migrate.go

# Тестування
test:
	go test ./...


#1lib офіційний драйвер PostgreSQL для стандартної бібліотеки database/sql у Go
#2 для роботи з ,l у стилі ORM, яка дозволяє писати SQL-запити через Go-структури.
#3 драйвер для GORM, який дозволяє GORM взаємодіяти з PostgreSQL.