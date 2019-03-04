CREATE TABLE users (
  id bigint NOT NULL,
  username    varchar(40) NOT NULL,
  email   varchar(40) NOT NULL,
  password   varchar(40) NOT NULL,
  CONSTRAINT users_id PRIMARY KEY (id)
);