-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN email_verified_at timestamp NULL DEFAULT NULL AFTER created_at;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN email_verified_at;
-- +goose StatementEnd
