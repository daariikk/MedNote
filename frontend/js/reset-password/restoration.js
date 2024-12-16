document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.querySelector('.login-form form');


    loginForm.addEventListener('submit', function (event) {
        event.preventDefault(); 


        const email = document.querySelector('input[type="email"]').value;
        const password = document.querySelector('input[type="password"]').value;
        const newPassword = document.querySelector('input[type="new-password"]').value;
        const repeatNewPassword = document.querySelector('input[type="repeat-new-password"]').value;

        if (newPassword !== repeatNewPassword) {
            alert('Пароли не совпадают');
            return;
        }

        // Создаем объект данных для отправки
        const loginData = {
            email: email,
            password: password,
            new_password: newPassword
        };

        // Отправляем POST-запрос на сервер
        fetch('http://localhost:8082/api/v1/auth/restoration', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(loginData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.status === 'success') {
                window.location.href = 'authorization.html';
            } else {
                alert('Ошибка авторизации: ' + (data.message || 'Неизвестная ошибка'));
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
            alert('Произошла ошибка при восстановлении пароля');
        });
    });
});