CREATE TABLE oauth_clients (
    id UUID PRIMARY KEY,
    client_id VARCHAR(100) UNIQUE NOT NULL,
    client_secret VARCHAR(255),
    name VARCHAR(100),
    redirect_uris TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE auth_code_sessions (
    code VARCHAR(255) PRIMARY KEY,
    client_id VARCHAR(100) NOT NULL,
    user_id UUID NOT NULL,
    code_challenge TEXT NOT NULL,
    redirect_uri TEXT NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    CONSTRAINT fk_client_id FOREIGN KEY (client_id) REFERENCES oauth_clients (client_id) ON DELETE CASCADE,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE INDEX idx_auth_code_sessions_expires_at ON auth_code_sessions (expires_at);
