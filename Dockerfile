# Используем официальный образ PostgreSQL
FROM postgres:latest

# Устанавливаем переменные окружения для PostgreSQL
# Устанавливаем переменные окружения для PostgreSQL
ENV POSTGRES_DB=med_note
ENV POSTGRES_USER=northwindman
ENV POSTGRES_PASSWORD=newpassword

# Открываем порт для подключения к базе данных
EXPOSE 5432