CREATE TABLE IF NOT EXISTS water(
    id BIGSERIAL PRIMARY KEY,
    volume_glass INT NOT NULL, 
    count_glass INT NOT NULL,
    indicator INT NOT NULL,                                             
    control VARCHAR(10) CHECK (control IN ('NORMAL', 'BAD', 'CRITICAL')) NOT NULL,  
    date_of_addition VARCHAR(10) NOT NULL,
    patient_id BIGINT NOT NULL,                                
    CONSTRAINT fk_water_patient FOREIGN KEY (patient_id) REFERENCES patients(patient_id) 
);