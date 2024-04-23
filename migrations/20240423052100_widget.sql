-- +goose Up
-- +goose StatementBegin
CREATE TABLE
  widgets (id serial PRIMARY KEY, color TEXT);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE widgets;

-- +goose StatementEnd