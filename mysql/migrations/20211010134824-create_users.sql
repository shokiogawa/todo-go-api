-- +migrate Up
CREATE TABLE IF NOT EXISTS users
(
    id             MEDIUMINT    NOT NULL AUTO_INCREMENT PRIMARY KEY,
    public_user_id VARCHAR(36)  NOT NULL,
    name           VARCHAR(32)  NOT NULL,
    email          VARCHAR(255) NOT NULL,
    password       VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE users
    ADD INDEX index_user (public_user_id);
-- +migrate Down
DROP TABLE IF EXISTS users;