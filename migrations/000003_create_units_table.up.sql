CREATE TABLE IF NOT EXISTS units (
    id UUID PRIMARY KEY,
    installation_id UUID NOT NULL,
    code VARCHAR(50) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_units_installation FOREIGN KEY (installation_id) REFERENCES installations(id)
);

CREATE INDEX idx_units_installation_id ON units(installation_id);
CREATE INDEX idx_units_deleted_at ON units(deleted_at);
