-- CREATE SEQUENCE id_seq;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
  id uuid NOT NULL DEFAULT uuid_generate_v4(),
  username    varchar(40) unique NOT NULL,
  email   varchar(40) unique NOT NULL,
  password   varchar(60) NOT NULL,
  role   varchar(40) NOT NULL,
  CONSTRAINT users_id PRIMARY KEY (id)
);