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
  FOREIGN KEY (technology_id) REFERENCES technology(id)
);

SELECT manage_updated_at('section');

CREATE TABLE task (
  id uuid NOT NULL,
  section_id uuid NOT NULL,
  title varchar(1024) NOT NULL,
  image_url varchar(1024),
  difficulty integer NOT NULL,
  content json NOT NULL,
  position integer NOT NULL,
  is_dynamic boolean NOT NULL DEFAULT FALSE,
  is_public boolean NOT NULL DEFAULT FALSE,
  updated_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  created_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  FOREIGN KEY (section_id) REFERENCES section(id)
);

SELECT manage_updated_at('task');
