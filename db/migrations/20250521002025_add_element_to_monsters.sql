-- +migrate Up
ALTER TABLE monster ADD COLUMN element VARCHAR(255) NULL;

-- +migrate Down
ALTER TABLE monster DROP COLUMN element;
