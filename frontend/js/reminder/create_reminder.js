document.addEventListener("DOMContentLoaded", function () {
    // Обработчик отправки формы
    document.getElementById("reminder-form").addEventListener("submit", function (event) {
        event.preventDefault(); // Предотвращаем стандартное поведение отправки формы

        const titleInput = document.querySelector(".title-input");
        const textInput = document.querySelector(".text-input");
        const datePicker = document.querySelector(".date-picker input");
        const timePicker = document.querySelector(".time-picker input");
        const patientId = parseInt(localStorage.getItem("patient_id"));



        // Проверка введённых данных
        if (!titleInput.value || !textInput.value || !datePicker.value || !timePicker.value) {
            alert("Пожалуйста, заполните все поля формы.");
            return;
        }

        // Создание объекта с данными для отправки
        const newReminder = {
            patient_id: patientId,
            title: titleInput.value,
            text: textInput.value,
            date: datePicker.value,
            time: timePicker.value + ":00"
        };

        console.log("newReminder: ", newReminder)

        // Отправка POST-запроса на сервер
        createReminder(newReminder)
            .then(data => {
                if (data.status === "success" && data.data) {
                    // Добавляем новое напоминание на страницу
                    const remindersContainer = document.querySelector(".reminders");
                    const content = `
                        <div class="reminder" data-reminder-id="${data.data.id}">
                            <div class="reminder-header">
                                <span>${formatDate(data.data.date)}</span>
                                <img class="calendar" src="../../img/all/calendar.png" alt="Calendar Icon">
                                <span>${formatTime(data.data.time)}</span>
                                <img class="clock" src="../../img/all/clock.png" alt="Time Icon">
                                <button class="delete-button">x</button>
                            </div>
                            <div class="reminder-content">
                                <h2>${data.data.title}</h2>
                                <p>${data.data.text}</p>
                            </div>
                        </div>
                    `;
                    remindersContainer.insertAdjacentHTML("beforeend", content);

                    // Очищаем форму
                    titleInput.value = "";
                    textInput.value = "";
                    datePicker.value = "";
                    timePicker.value = "";

                    console.log("Напоминание успешно добавлено:", data.data);
                } else {
                    console.error("Ошибка при добавлении напоминания:", data);
                }
            })
            .catch(error => {
                console.error("Ошибка при отправке запроса:", error);
            });
    });
    // Функция для форматирования даты
    function formatDate(dateString) {
        const date = new Date(dateString);
        const day = date.getDate().toString().padStart(2, '0');
        const month = (date.getMonth() + 1).toString().padStart(2, '0');
        const year = date.getFullYear();
        return `${day}.${month}.${year}`;
    }

    // Функция для форматирования времени
    function formatTime(timeString) {
        const [hours, minutes] = timeString.split(":");
        return `${hours}:${minutes}`;
    }
    // Функция для создания напоминания
    function createReminder(reminderData) {
        const createUrl = `http://localhost:8082/api/v1/reminders`;
        return fetch(createUrl, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${localStorage.getItem('token')}` 
            },
            body: JSON.stringify(reminderData)
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Ошибка при создании напоминания");
                }
                return response.json();
            });
    }
});