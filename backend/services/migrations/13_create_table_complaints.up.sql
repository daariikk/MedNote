CREATE TABLE IF NOT EXISTS complaints (
    id BIGSERIAL PRIMARY KEY,
    complaint VARCHAR(512) NOT NULL,                                             
    date_of_addition VARCHAR(10) NOT NULL,
    patient_id BIGINT NOT NULL,                                
    CONSTRAINT fk_complaints_patient FOREIGN KEY (patient_id) REFERENCES patients(patient_id) 
);