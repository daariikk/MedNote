# MedNote — Удобное ведение медицинских записей

![MedNote Logo](frontend/img/all/logo.png)

MedNote — это приложение, которое помогает вам вести медицинские записи, следить за вашим здоровьем и вводить осознанность в свою жизнь. С помощью MedNote вы можете легко отслеживать свои показатели, ставить напоминания о важных событиях и делиться отчетами с врачами.

---

## Основные функции

### 1. **Ведение медицинских записей**
Записывайте ваши показатели здоровья и отслеживайте их изменения. Приложение подсветит отклонения, на которые стоит обратить внимание.

### 2. **Напоминания о важном**
Добавляйте напоминания о приеме лекарств, посещении врача или тренировках. Забота о здоровье начинается с полезных привычек!

---

## Технологии

- **Backend**: Go, PostgreSQL, Docker, RabbitMQ
- **Frontend**: HTML, CSS, JavaScript
- **Инструменты**: Taskfile для автоматизации задач

---

## Установка и запуск

### Предварительные требования
- Убедитесь, что у вас установлены:
  - [Go](https://golang.org/dl/)
  - [Docker](https://www.docker.com/products/docker-desktop)
  - [PostgreSQL](https://www.postgresql.org/download/)

### Шаги для запуска

1. **Клонируйте репозиторий:**

   ```
   git clone https://github.com/daariikk/MedNote.git
   cd MedNote
   ```

3. **Скопируйте файл .env.example в .env и настройте переменные окружения.**

4. **Запуск миграций:**
  ```
  task migrate_up
  ```

4. **Запуск Docker контейнеров:**
  ```
  docker-compose up
  ```

5. **Запуск сервисов:**
  ```
  task run_api_gateway
  task run_patient
  task run_reminder
  task run_record
  task run_notification
  ```

6. **Запуск фронтенда:**

Откройте файл frontend/index.html в вашем браузере.

---

## Команда
Шкарупа Дарья Евгеньевна - System Analyst, Backend Developer, Frontend Developer, Design
