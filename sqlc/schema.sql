CREATE TYPE user_status AS ENUM ('active', 'inactive');
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    status user_status DEFAULT 'active',
    role_id INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_role FOREIGN KEY (role_id) REFERENCES roles(id)
);
CREATE TABLE keys (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    refresh_token TEXT,
    used_refresh_token TEXT[],
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT unique_user_id UNIQUE (user_id)
);
