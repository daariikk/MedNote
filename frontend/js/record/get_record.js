document.addEventListener('DOMContentLoaded', function () {
const today = new Date().toISOString().split('T')[0];

// Находим поле ввода даты
const dateInput = document.getElementById('custom-date');

// Устанавливаем значение поля в текущую дату
dateInput.value = today;
sendGetRequest(today);

dateInput.addEventListener('change', function () {
    const selectedDate = dateInput.value; // Получаем выбранную дату
    sendGetRequest(selectedDate);
});
});

// Функция для отправки GET-запроса
function sendGetRequest(date) {
    const patientId = parseInt(localStorage.getItem("patient_id")) || 1;

    // Проверяем, что patient_id существует
    if (!patientId) {
        console.error('patient_id не найден в localStorage');
        return;
    }

    // Отправляем GET-запрос на сервер
    fetch(`http://localhost:8082/api/v1/records?userId=${patientId}&date=${date}`,
        {
            method: "GET",
            headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}` // Добавляем токен
        }
        }
    )
        .then(response => {
            if (response.ok) {
                return response.json(); // Преобразуем ответ в JSON
            } else {
                throw new Error('Ошибка при получении данных: ' + response.status);
            }
        })
        .then(data => {
            console.log('Ответ сервера:', data);

            // Очищаем контейнер перед добавлением новых карточек
            const recordsContainer = document.querySelector('.records');
            recordsContainer.innerHTML = '';

            // Отрисовываем каждую запись
            data.data.forEach(record => {
                const recordData = getRecordTemplate(record); // Получаем шаблон для записи
                const recordElement = createRecordElement(recordData, record.type, record.id, record.control); // Создаем элемент записи
                recordsContainer.appendChild(recordElement); // Добавляем элемент в контейнер
            });
        })
        .catch(error => {
            console.error('Ошибка:', error);
        });
}