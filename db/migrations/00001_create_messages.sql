-- +goose Up
-- SQL in this section is executed when the migrations is applied.
CREATE TABLE IF NOT EXISTS messages (
    id                    BIGINT AUTO_INCREMENT,
    content               TEXT NOT NULL,
    created_at            DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at            DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
    );

-- +goose Down
-- SQL in this section is executed when the migrations is rolled back.
DROP TABLE IF EXISTS messages;