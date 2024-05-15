-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  magic_link (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users (id) ON DELETE CASCADE,
    ml_token_hash TEXT UNIQUE NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL
  );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE magic_link;
-- +goose StatementEnd