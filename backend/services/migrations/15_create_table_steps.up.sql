CREATE TABLE IF NOT EXISTS steps (
    id BIGSERIAL PRIMARY KEY,
    indicator INT NOT NULL,                                              
    control VARCHAR(10) CHECK (control IN ('NORMAL', 'BAD', 'CRITICAL')) NOT NULL,  
    date_of_addition VARCHAR(10) NOT NULL,
    patient_id BIGINT NOT NULL,                                
    CONSTRAINT fk_steps_patient FOREIGN KEY (patient_id) REFERENCES patients(patient_id) 
);