-- +goose Up

CREATE TABLE IF NOT EXISTS blogs(
    id UUID PRIMARY KEY NOT NULL,
    userId UUID NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    content TEXT NOT NULL,
    thumbnail TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(userId) REFERENCES users(id)
);

-- +goose Down

DROP TABLE IF EXISTS blogs
