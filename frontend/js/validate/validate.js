// Настройка валидации для полей
function setupValidation() {
    const editableFields = document.querySelectorAll(".editable");

    editableFields.forEach(field => {
        field.addEventListener("input", function () {
            const fieldType = field.getAttribute("data-field");
            switch (fieldType) {
                case "name":
                    validateName(field);
                    break;
                case "surname":
                    validateSurname(field);
                    break;
                case "gender":
                    validateGender(field);
                    break;
            }
        });
    });
}

// Валидация имени
function validateName(field) {
    field.textContent = field.textContent.replace(/[^А-Яа-яA-Za-z\s]/g, ''); // Удаляем все, кроме букв и пробелов
}

// Валидация фамилии
function validateSurname(field) {
    field.textContent = field.textContent.replace(/[^А-Яа-яA-Za-z\s]/g, ''); // Удаляем все, кроме букв и пробелов
}

// Валидация пола
function validateGender(field) {
    const allowedGenders = ['М', 'Ж'];
    if (!allowedGenders.includes(field.textContent)) {
        field.textContent = ''; // Очищаем поле, если ввод некорректен
    }
}
