document.addEventListener("DOMContentLoaded", function () {
    const patientId = localStorage.getItem("patient_id");
    const userName = localStorage.getItem("user_name");

    const logoLink = document.getElementById("logo-link");

    const authButton = document.querySelector(".auth-buttons button");


    if (patientId) {
        logoLink.href = "../records/record.html";

        if (userName) {
            authButton.href = "../homepage/account.html"
            authButton.textContent = userName;
        } else {
            authButton.href = "../homepage/account.html"
            authButton.textContent = "Мой аккаунт"; 
        }
    } else {
        logoLink.href = "../homepage/index.html"
        authButton.href = "../auth/authorization.html";
        authButton.textContent = "Войти";
    }
});

window.addEventListener("storage", function (event) {
    if (event.key === "patient_id") {
        const logoLink = document.getElementById("logo-link");
        const authButton = document.querySelector(".auth-buttons button");
        const userName = localStorage.getItem("user_name");

        if (event.newValue) {
            logoLink.href = "../homepage/account.html";
            if (userName) {
                authButton.textContent = userName;
            } else {
                authButton.textContent = "Мой аккаунт"; 
            }
        } else {
            logoLink.href = "../homepage/index.html";
            authButton.textContent = "Войти";
        }
    }
});


