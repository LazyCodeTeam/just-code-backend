BEGIN;

ALTER TABLE task
ALTER COLUMN title
DROP NOT NULL;

COMMIT;