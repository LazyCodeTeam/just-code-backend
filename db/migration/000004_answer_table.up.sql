BEGIN;

CREATE TYPE answer_result AS ENUM (
  'FIRST_VALID', 
  'VALID', 
  'INVALID'
);

CREATE TABLE answer (
  id BIGSERIAL,
  task_id uuid NOT NULL,
  profile_id varchar(64) NOT NULL,
  result answer_result NOT NULL,
  created_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  FOREIGN KEY (task_id) REFERENCES task (id) ON DELETE CASCADE
);

CREATE INDEX answer_task_id_idx ON answer (task_id);
CREATE INDEX answer_created_at_idx ON answer (created_at);
CREATE INDEX answer_result_idx ON answer (result);
CREATE INDEX answer_profile_id_idx ON answer (profile_id);

COMMIT;
