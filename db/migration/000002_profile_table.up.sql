CREATE TABLE profile (
  id varchar(64) NOT NULL,
  name varchar(64) NOT NULL UNIQUE,
  avatar_url varchar(1024),
  first_name varchar(64),
  last_name varchar(64),
  updated_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  created_at timestamp with TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id)
);

SELECT manage_updated_at('profile');
