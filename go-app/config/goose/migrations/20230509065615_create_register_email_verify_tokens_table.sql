-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `register_email_verify_tokens` (
  `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
  `user_id` VARCHAR(255) NOT NULL,
  `email` VARBINARY(255) NOT NULL,
  `token` VARCHAR(255) NOT NULL UNIQUE,
  `expired_at` DATETIME NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME
)
ENGINE = InnoDB
COMMENT = '登録前のメールアドレス検証用トークン';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `register_email_verify_tokens`;
-- +goose StatementEnd
