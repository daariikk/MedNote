document.addEventListener('DOMContentLoaded', function () {
    const loginForm = document.querySelector('.login-form form');

    loginForm.addEventListener('submit', function (event) {
        event.preventDefault(); 


        const email = document.querySelector('input[type="email"]').value;
        const password = document.querySelector('input[type="password"]').value;

        // Создаем объект данных для отправки
        const loginData = {
            email: email,
            password: password
        };

        // Отправляем POST-запрос на сервер
        fetch('http://localhost:8082/api/v1/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(loginData)
        })
        .then(response => response.json())
        .then(data => {
            if (data.status === 'success') {
                localStorage.setItem('token', data.data.token);
                localStorage.setItem('patient_id', data.data.patient_id);

                window.location.href = '../homepage/account.html';
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