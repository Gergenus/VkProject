-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL -- bcrypt hash
);

CREATE INDEX IF NOT EXISTS idx_login ON users(login);

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    subject TEXT NOT NULL,
    post_text TEXT NOT NULL,
    image_address TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts(user_id);  
CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_posts_created_at;
DROP INDEX IF EXISTS idx_posts_user_id;
DROP TABLE IF EXISTS posts;
DROP INDEX IF EXISTS idx_login;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
