function updateCardContent() {
    const recordType = document.getElementById('record-type').value;
    const cardContent = document.getElementById('card-content');

    let content = '';

    if (recordType === 'pulse') {
        content = `
            <div class="line-1">
            <label for="pulse">Пульс:</label>
            <input type="number" id="pulse" placeholder="__ уд. в мин.">
            </div>
        `;
    } else if (recordType === 'tensions') {
        content = `
            <div class="line-1">
            <label for="upper-pressure">Верхнее давление:</label>
            <input type="number" id="upper-pressure" placeholder="__ мм. Рт. Ст.">
            </div>
            <div class="line-2">
            <label for="lower-pressure">Нижнее давление:</label>
            <input type="number" id="lower-pressure" placeholder="__ мм. Рт. Ст.">
            </div>
        `;
    } else if (recordType === 'steps') {
        content = `
            <div class="line-1">
            <label for="steps">Количество шагов:</label>
            <input type="number" id="steps" placeholder="__ шагов">
            </div>
        `;
    } else if (recordType === 'sleep') {
        content = `
            
            <label for="sleep-duration">Длительность сна:</label>
            <div class="line-1">
            <input type="number" id="sleep-hours" placeholder="__ h"> часов
            <input type="number" id="sleep-minutes" placeholder="__ m"> минут
            </div>
        `;
    } else if (recordType === 'water') {
        content = `
            <div class="line-1">
            <label for="water">Объем стакана:</label>
            <input type="number" id="water" placeholder="__ мл">
            </div>
            <div class="line-2">
            <label for="number">Количество выпитых стаканов:</label>
            <input type="number" id="number" placeholder="__ шт">
            </div>
        `;
    } else if (recordType === 'complaints') {
        content = `
            <div class="line-complaints">
            <label for="complaints">Жалоба:</label>
            <textarea id="complaints" placeholder="Введите вашу жалобу"></textarea>
            </div>
        `;
    }

    cardContent.innerHTML = content;
}



// Функция для добавления новой записи
function addRecord(event) {
    event.preventDefault(); // Предотвращаем перезагрузку страницы при отправке формы

    const recordType = document.getElementById('record-type').value;
    const recordsContainer = document.querySelector('.records');
    const date = document.getElementById("custom-date").value;
    const patientId = parseInt(localStorage.getItem("patient_id"));


    let recordData = '';
    if (recordType === 'pulse') {
        const pulseValue = document.getElementById('pulse').value;
        const control = pulseControl(pulseValue);

        record = {
            type: recordType,
            indicator: parseInt(pulseValue, 10),
            control: control,
            date_of_addition: date,
            patient_id: patientId
        };

        // Отправляем запись на сервер
        sendRecordToServer(record, recordsContainer);

    } else if (recordType === 'tensions') {
        const upperPressure = parseInt(document.getElementById('upper-pressure').value, 10);
        const lowerPressure = parseInt(document.getElementById('lower-pressure').value, 10);
        const control = tensionsControl(upperPressure, lowerPressure);
        record = {
            type: recordType,
            upper_indicator: upperPressure,
            lower_indicator: lowerPressure,
            control: control,
            date_of_addition: date,
            patient_id: patientId
        };

        // Отправляем запись на сервер
        sendRecordToServer(record, recordsContainer);

    } else if (recordType === 'steps') {
        const steps = parseInt(document.getElementById('steps').value, 10);
        const control = stepsControl(steps);
        record = {
            type: recordType,
            indicator: steps,
            control: control,
            date_of_addition: date,
            patient_id: patientId
        };
        sendRecordToServer(record, recordsContainer);
    } else if (recordType === 'sleep') {
        const sleepHours = parseInt(document.getElementById('sleep-hours').value, 10);
        const sleepMinutes = parseInt(document.getElementById('sleep-minutes').value, 10);
        const control = sleepControl(sleepHours, sleepMinutes);
        record = {
            type: recordType,
            hours: sleepHours,
            minutes: sleepMinutes,
            control: control,
            date_of_addition: date,
            patient_id: patientId
        };
        sendRecordToServer(record, recordsContainer);

    } else if (recordType === 'water') {
        const waterVolume = parseInt(document.getElementById('water').value, 10);
        const glassesCount = parseInt(document.getElementById('number').value, 10);
        const control = waterControl(waterVolume, glassesCount);

        record = {
            type: recordType,
            volume_glass: waterVolume,
            count_glass: glassesCount,
            indicator: waterVolume*glassesCount,
            control: control,
            date_of_addition: date,
            patient_id: patientId
        };
        sendRecordToServer(record, recordsContainer);

    } else if (recordType === 'complaints') {
        const complaint = document.getElementById('complaints').value;
        record = {
            type: recordType,
            complaint: complaint,
            date_of_addition: date,
            patient_id: patientId
        };
        sendRecordToServer(record, recordsContainer);
    }

    // Создаем элемент записи
    const recordElement = document.createElement('div');
    recordElement.classList.add('record-card');
    recordElement.setAttribute('data-id', newRecord.data.id);
    recordElement.innerHTML = `
        <p>${recordData}</p>
        <button class="delete-btn" onclick="deleteRecord(this)">&#10006;</button>
    `;

    // Добавляем новую запись в контейнер
    recordsContainer.appendChild(recordElement);
}


// Функция для отправки записи на сервер
function sendRecordToServer(record, recordsContainer) {
    console.log("Сформированный JSON:", JSON.stringify(record, null, 2));

    fetch('http://localhost:8082/api/v1/records', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            "Authorization": `Bearer ${localStorage.getItem('token')}`
        },
        body: JSON.stringify(record)
    })
        .then(response => {
            if (response.ok) {
                return response.json();
            } else {
                throw new Error('Ошибка при отправке запроса: ' + response.status);
            }
        })
        .then(data => {
            console.log('Ответ сервера:', data);
            alert('Данные успешно отправлены!');


            const recordData = getRecordTemplate(record);
            const recordElement = createRecordElement(recordData, record.type, data.data.id, record.control);

            recordsContainer.appendChild(recordElement);
        })
        .catch(error => {
            console.error('Ошибка:', error);
            alert('Произошла ошибка при отправке данных.');
        });
}


function getRecordTemplate(record) {
    switch (record.type) {
        case 'pulse':
            return `Пульс: ${record.indicator} уд. в мин.`;
        case 'tensions':
            return `Давление: ${record.upper_indicator}/${record.lower_indicator} мм. Рт. Ст.`;
        case 'steps':
            return `Шаги: ${record.indicator}`;
        case 'sleep':
            return `Сон: ${record.hours} ч. ${record.minutes} мин.`;
        case 'water':
            return `Объём выпитой воды: ${record.volume_glass * record.count_glass} мл`;
        case 'complaints':
            return `Жалоба: ${record.complaint}`;
        default:
            return 'Неизвестный тип записи';
    }
}

// Функция для создания элемента записи
function createRecordElement(recordData, type=null, recordId = null, control = null) {
    const recordElement = document.createElement('div');
    recordElement.classList.add('record-card');

    // Добавляем класс в зависимости от значения control
    if (control) {
        recordElement.classList.add(`control-${control}`);
    }

    if (recordId) {
        recordElement.setAttribute('data-id', recordId); // Добавляем id, если он есть
    }

    if (type) {
        recordElement.setAttribute('data-type', type);
    }

    recordElement.innerHTML = `
        <p>${recordData}</p>
        <button class="delete-btn" onclick="deleteRecord(this)">&#10006;</button>
    `;

    return recordElement;
}

// Функция для удаления записи
function deleteRecord(button) {
    // Находим родительский элемент (record-card)
    const recordCard = button.closest('.record-card');
    const patientId = parseInt(localStorage.getItem("patient_id"));
    const tableName = recordCard.getAttribute('data-type');
    const recordId = parseInt(recordCard.getAttribute('data-id'));

    console.log("patientId: ", patientId)
    console.log("tableName: ", tableName)
    console.log("recordId: ", recordId)
    // Проверяем, что все данные получены
    if (!patientId || !tableName || !recordId) {
        console.error('Не удалось получить данные для удаления записи.');
        return;
    }

    const url = `http://localhost:8082/api/v1/records?userId=${patientId}&tableName=${tableName}&recordId=${recordId}`
    console.log("url: ", url)
    // Отправляем DELETE-запрос на сервер
    fetch(`http://localhost:8082/api/v1/records?userId=${patientId}&tableName=${tableName}&recordId=${recordId}`, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
            "Authorization": `Bearer ${localStorage.getItem('token')}`
        }
    })
    .then(response => {
        if (response.ok) {
            // Удаляем карточку из DOM
            recordCard.remove();
            console.log(`Запись с ID ${recordId} успешно удалена.`);
        } else {
            console.error('Ошибка при удалении записи:', response.status);
        }
    })
    .catch(error => {
        console.error('Ошибка:', error);
    });
}
