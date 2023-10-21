CREATE TABLE users(
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    last_login_attempt TIMESTAMP
);

CREATE UNIQUE INDEX users_useraname ON users(username);

CREATE TABLE user_sessions(
    id TEXT PRIMARY KEY,
    user_id UUID NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
