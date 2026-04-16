CREATE TABLE IF NOT EXISTS specialties (
    id UUID PRIMARY KEY,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_specialties_deleted_at ON specialties(deleted_at);

CREATE TABLE IF NOT EXISTS doctors (
    id UUID PRIMARY KEY,
    specialty_id UUID NOT NULL,
    nik VARCHAR(50) UNIQUE NOT NULL,
    sip VARCHAR(100) UNIQUE NOT NULL,
    sip_expiry_date DATE NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_doctors_specialty FOREIGN KEY (specialty_id) REFERENCES specialties(id)
);

CREATE INDEX idx_doctors_specialty_id ON doctors(specialty_id);
CREATE INDEX idx_doctors_deleted_at ON doctors(deleted_at);

CREATE TABLE IF NOT EXISTS doctor_units (
    doctor_id UUID NOT NULL,
    unit_id UUID NOT NULL,
    PRIMARY KEY (doctor_id, unit_id),
    CONSTRAINT fk_doctor_units_doctor FOREIGN KEY (doctor_id) REFERENCES doctors(id) ON DELETE CASCADE,
    CONSTRAINT fk_doctor_units_unit FOREIGN KEY (unit_id) REFERENCES units(id) ON DELETE CASCADE
);
