document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.querySelector('.login-form form');

    loginForm.addEventListener('submit', function (event) {
        event.preventDefault(); 


        const email = document.querySelector('input[type="email"]').value;

        // Создаем объект данных для отправки
        const loginData = {
            email: email,
        };

        // Отправляем POST-запрос на сервер
        fetch('http://localhost:8082/api/v1/auth/reset-password', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(loginData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.status === 'success') {
                alert('Письмо отправлено на почту! Действуйте по интсрукции в письме.');
            } else {
                alert('Ошибка авторизации: ' + (data.message || 'Неизвестная ошибка'));
            }
        })
        .catch(error => {
            console.error('Ошибка:', error);
            alert('Произошла ошибка при авторизации');
        });
    });
});