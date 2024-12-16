document.addEventListener("DOMContentLoaded", function () {
    const userId = localStorage.getItem('patient_id');
    if (!userId) {
        console.error("Ошибка: patient_id не найден в localStorage");
        return;
    }
    const apiUrl = `http://localhost:8082/api/v1/users/${userId}/patient`;

    // Отправка GET-запроса на сервер
    fetch(apiUrl, {
        method: "GET",
        headers: {
            "Authorization": `Bearer ${localStorage.getItem('token')}` 
        }
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Ошибка сети или сервера");
            }
            return response.json();
        })
        .then(data => {
            if (data.status === "success" && data.data) {
                localStorage.setItem('user_name', data.data.first_name);

                fillUserData(data.data);
            } else {
                console.error("Некорректный ответ от сервера:", data);
            }
        })
        .catch(error => {
            console.error("Произошла ошибка:", error);
        });

        setupValidation();
});

// Функция для заполнения данных на странице
function fillUserData(patientData) {
    document.querySelector('[data-field="name"]').textContent = patientData.first_name;
    document.querySelector('[data-field="surname"]').textContent = patientData.second_name;
    document.querySelector('[data-field="email"]').textContent = patientData.email;
    document.querySelector('[data-field="gender"]').textContent = patientData.gender;
    document.querySelector('[data-field="height"]').textContent = patientData.height;
    document.querySelector('[data-field="weight"]').textContent = patientData.weight;
}

// Функция для переключения режима редактирования
function toggleEditMode() {
    const editableFields = document.querySelectorAll(".editable");
    const editButton = document.getElementById("editButton");
    const saveButton = document.getElementById("saveButton");

    // Переключение режима редактирования
    editableFields.forEach(field => {
        if (field.contentEditable === "true") {
            field.contentEditable = "false";
        } else {
            field.contentEditable = "true";
        }
    });

    // Показ/скрытие кнопок
    editButton.style.display = editButton.style.display === "none" ? "inline" : "none";
    saveButton.style.display = saveButton.style.display === "none" ? "inline" : "none";
}

// Функция для сохранения изменений
function saveChanges() {
    const userId = localStorage.getItem('patient_id');
    if (!userId) {
        console.error("Ошибка: patient_id не найден в localStorage");
        return;
    }

    const apiUrl = `http://localhost:8082/api/v1/users/${userId}/patient`;

    // Сбор данных из полей
    const updatedData = {
        first_name: document.querySelector('[data-field="name"]').textContent,
        second_name: document.querySelector('[data-field="surname"]').textContent,
        height: parseFloat(document.querySelector('[data-field="height"]').textContent),
        weight: parseFloat(document.querySelector('[data-field="weight"]').textContent),
        gender: document.querySelector('[data-field="gender"]').textContent
    };

    // Отправка PUT-запроса на сервер
    fetch(apiUrl, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${localStorage.getItem('token')}`
            
        },
        body: JSON.stringify(updatedData)
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Ошибка сети или сервера");
            }
            return response.json();
        })
        .then(data => {
            if (data.status === "success" && data.data) {
                fillUserData(data.data);

                toggleEditMode();
            } else {
                console.error("Некорректный ответ от сервера:", data);
            }
        })
        .catch(error => {
            console.error("Произошла ошибка:", error);
        });
}