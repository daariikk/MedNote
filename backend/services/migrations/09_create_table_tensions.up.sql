CREATE TABLE IF NOT EXISTS tensions (
    id BIGSERIAL PRIMARY KEY,
    upper_indicator INT NOT NULL,
    lower_indicator INT NOT NULL,
    control VARCHAR(10) CHECK (control IN ('NORMAL', 'BAD', 'CRITICAL')) NOT NULL,  
    date_of_addition VARCHAR(10) NOT NULL,
    patient_id BIGINT NOT NULL,                                
    CONSTRAINT fk_tensions_patient FOREIGN KEY (patient_id) REFERENCES patients(patient_id) 
);