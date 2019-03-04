CREATE SEQUENCE id_seq;

CREATE TABLE users (
  id bigint NOT NULL DEFAULT nextval('id_seq'),
  username    varchar(40) unique NOT NULL,
  email   varchar(40) unique NOT NULL,
  password   varchar(40) NOT NULL,
  role   varchar(40) NOT NULL,
  CONSTRAINT users_id PRIMARY KEY (id)
);