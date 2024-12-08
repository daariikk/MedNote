CREATE TABLE IF NOT EXISTS pulse (
    id BIGSERIAL  PRIMARY KEY,
    indicator VARCHAR(4) NOT NULL,                                              
    control VARCHAR(10) CHECK (control IN ('NORMAL', 'BAD', 'CRITICAL')) NOT NULL,  
    date_of_addition DATE NOT NULL,                            
    patient_id BIGINT NOT NULL,                                
    CONSTRAINT fk_pulse_patient FOREIGN KEY (patient_id) REFERENCES patients(patient_id) 
);