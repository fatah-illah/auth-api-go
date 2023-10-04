-- Users Table
CREATE TABLE public.users (
    id serial PRIMARY KEY,
    username varchar(255) UNIQUE NOT NULL,
    password_hash varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
SELECT * FROM users;

-- Tokens Table
CREATE TABLE public.tokens (
    id serial PRIMARY KEY,
    user_id integer REFERENCES users(id) ON DELETE CASCADE,
    access_token varchar(255) UNIQUE NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    expires_at timestamp without time zone
);
SELECT * FROM tokens;

-- Protected Resources Table
CREATE TABLE public.protected_resources (
    resource_id serial PRIMARY KEY,
    resource_name varchar(255) NOT NULL,
    resource_data text NOT NULL
);
SELECT * FROM protected_resources;

-- Insert Dummy Users
INSERT INTO users (username, password_hash, email) VALUES 
('john_doe', 'hashed_password_1', 'john_doe@example.com'),
('jane_smith', 'hashed_password_2', 'jane_smith@example.com'),
('alice_wong', 'hashed_password_3', 'alice_wong@example.com');

SELECT * FROM users;

-- Insert Dummy Tokens (Harap dicatat bahwa ini hanya contoh; dalam praktiknya, Anda harus menghasilkan token yang unik dan aman)
INSERT INTO tokens (user_id, access_token, expires_at) VALUES 
(1, 'token_1_for_john', '2024-01-01 23:59:59'),
(2, 'token_2_for_jane', '2024-01-01 23:59:59'),
(3, 'token_3_for_alice', '2024-01-01 23:59:59');

SELECT * FROM tokens;

-- Insert Dummy Protected Resources
INSERT INTO protected_resources (resource_name, resource_data) VALUES 
('Resource_1', 'This is a protected data for Resource_1'),
('Resource_2', 'This is a protected data for Resource_2'),
('Resource_3', 'This is a protected data for Resource_3');

SELECT * FROM protected_resources;
