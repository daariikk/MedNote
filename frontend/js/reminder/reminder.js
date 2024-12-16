document.addEventListener("DOMContentLoaded", function () {
    // Замените на реальный userId
    const userId = localStorage.getItem('patient_id');
    if (!userId) {
        console.error("Ошибка: patient_id не найден в localStorage");
        return;
    }

    const apiUrl = `http://localhost:8082/api/v1/reminders?userId=${userId}`;

    // Функция для отрисовки карточек напоминаний
    function renderReminders(reminders) {
        const remindersContainer = document.querySelector(".reminders");
        remindersContainer.innerHTML = ""; // Очищаем контейнер перед отрисовкой

        reminders.forEach(reminder => {
            const content = `
                <div class="reminder" data-reminder-id="${reminder.id}">
                    <div class="reminder-header">
                        <span>${formatDate(reminder.date)}</span>
                        <img class="calendar" src="../../img/all/calendar.png" alt="Calendar Icon">
                        <span>${formatTime(reminder.time)}</span>
                        <img class="clock" src="../../img/all/clock.png" alt="Time Icon">
                        <button class="delete-button">x</button>
                    </div>
                    <div class="reminder-content">
                        <h2>${reminder.title}</h2>
                        <p>${reminder.text}</p>
                    </div>
                </div>
            `;
            remindersContainer.insertAdjacentHTML("beforeend", content);
        });
    }

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

    console.log("apiUrl: ", apiUrl)
    console.log("userId: " , userId)
    console.log("token: ", localStorage.getItem('token'))
    // Отправка GET-запроса на сервер
    fetch(apiUrl,
        {
            method: "GET",
            headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}` // Добавляем токен
        }
        }

)
        .then(response => {
            if (!response.ok) {
                throw new Error("Ошибка сети или сервера");
            }
            return response.json();
        })
        .then(data => {
            if (data.status === "success" && data.data) {
                renderReminders(data.data); // Отрисовка карточек
            } else {
                console.error("Некорректный ответ от сервера:", data);
            }
        })
        .catch(error => {
            console.error("Произошла ошибка:", error);
        });

         // Обработчик клика на кнопку удаления
    document.addEventListener("click", function (event) {
        if (event.target.classList.contains("delete-button")) {
            const reminderCard = event.target.closest(".reminder");
            const reminderId = reminderCard.getAttribute("data-reminder-id");

            // Подтверждение удаления (опционально)
            if (!confirm("Вы уверены, что хотите удалить это напоминание?")) {
                return;
            }

            // Отправка DELETE-запроса на сервер
            deleteReminder(userId, reminderId)
                .then(response => {
                    if (response.ok) {
                        // Удаляем карточку из DOM
                        reminderCard.remove();
                        console.log(`Напоминание с ID ${reminderId} успешно удалено.`);
                    } else {
                        console.error('Ошибка при удалении напоминания:', response.status);
                    }
                })
                .catch(error => {
                    console.error('Ошибка:', error);
                });
    
        }
    });


    // Функция для удаления напоминания
    function deleteReminder(userId, reminderId) {
        const deleteUrl = `http://localhost:8082/api/v1/reminders?userId=${userId}&reminderId=${reminderId}`;
        return fetch(deleteUrl,{
            method: "DELETE",
            headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}` // Добавляем токен
        }
})
}
})
;