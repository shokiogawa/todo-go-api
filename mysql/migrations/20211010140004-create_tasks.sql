
-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks
(
    id             MEDIUMINT    NOT NULL AUTO_INCREMENT PRIMARY KEY,
    public_task_id VARCHAR(255)  NOT NULL,
    user_id        MEDIUMINT    NOT NULL,
    title          VARCHAR(255) NOT NULL,
    content        VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
    );
ALTER TABLE tasks
    ADD INDEX index_task (public_task_id);
-- +migrate Down
DROP TABLE IF EXISTS tasks;