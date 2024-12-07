CREATE TABLE IF NOT EXISTS reminders (
    reminder_id BIGINT PRIMARY KEY,                         
    title VARCHAR(256) NOT NULL,                             
    text TEXT NOT NULL,                                      
    date DATE NOT NULL,                                     
    time TIME NOT NULL,                                     
    patient_id BIGINT NOT NULL,                            
    CONSTRAINT fk_reminders_patient FOREIGN KEY (patient_id) REFERENCES patients(patient_id) 
);