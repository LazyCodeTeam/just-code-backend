BEGIN;

CREATE TABLE technology (
  id uuid NOT NULL,
  title varchar(1024) NOT NULL,
  description text,
  image_url varchar(1024),
  position integer NOT NULL,
  updated_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  created_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

SELECT manage_updated_at('technology');


CREATE TABLE section (
  id uuid NOT NULL,
  technology_id uuid NOT NULL,
  title varchar(1024) NOT NULL,
  description text,
  image_url varchar(1024),
  position integer NOT NULL,
  updated_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  created_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  FOREIGN KEY (technology_id) REFERENCES technology(id) ON DELETE CASCADE
);

SELECT manage_updated_at('section');
CREATE INDEX section_technology_id_idx ON section (technology_id);

CREATE TABLE task (
  id uuid NOT NULL,
  section_id uuid NOT NULL,
  title varchar(1024) NOT NULL,
  description text,
  image_url varchar(1024),
  difficulty integer NOT NULL,
  content json NOT NULL,
  position integer,
  is_public boolean NOT NULL DEFAULT FALSE,
  updated_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  created_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  FOREIGN KEY (section_id) REFERENCES section(id) ON DELETE CASCADE
);

SELECT manage_updated_at('task');
CREATE INDEX task_section_id_idx ON task (section_id);

COMMIT;
