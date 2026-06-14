CREATE DATABASE grafana;
CREATE DATABASE "ssmDB";

\c grafana
CREATE TABLE IF NOT EXISTS "user" (
    id serial PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO "user" (login, email) VALUES ('admin', 'admin@example.com');

\c "ssmDB"
CREATE TABLE IF NOT EXISTS settings (
    id serial PRIMARY KEY,
    key VARCHAR(255) NOT NULL,
    value TEXT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO settings (key, value) VALUES ('version', '1.0.0');
