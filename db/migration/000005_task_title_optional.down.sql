BEGIN;

ALTER TABLE task
ALTER COLUMN title
SET DEFAULT '';

ALTER TABLE task
ALTER COLUMN title
Set
  NOT NULL;

COMMIT;
