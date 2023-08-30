BEGIN;

DROP INDEX IF EXISTS task_section_id_idx;
DROP INDEX IF EXISTS section_technology_id_idx;

DROP TABLE IF EXISTS task cascade;
DROP TABLE IF EXISTS section cascade;
DROP TABLE IF EXISTS technology cascade;
DROP TABLE IF EXISTS asset cascade;

COMMIT;
