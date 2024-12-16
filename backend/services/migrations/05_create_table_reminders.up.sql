CREATE TABLE IF NOT EXISTS reminders (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(256) NOT NULL,                             
    text TEXT NOT NULL,                                      
    date VARCHAR(10) NOT NULL,
    time VARCHAR(10) NOT NULL,
    patient_id BIGINT NOT NULL,                            
    CONSTRAINT fk_reminders_patient FOREIGN KEY (patient_id) REFERENCES patients(patient_id) 
);