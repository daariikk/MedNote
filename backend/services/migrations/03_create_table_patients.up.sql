CREATE TABLE IF NOT EXISTS patients (
    patient_id BIGSERIAL PRIMARY KEY, 
    first_name VARCHAR(100) NOT NULL, 
    second_name VARCHAR(100), 
    email VARCHAR(100) UNIQUE NOT NULL, 
    height FLOAT, 
    weight FLOAT, 
    gender CHAR(1) CHECK (gender IN ('Ж', 'М', NULL)), 
    password VARCHAR(100) NOT NULL, 
    registration_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
);
