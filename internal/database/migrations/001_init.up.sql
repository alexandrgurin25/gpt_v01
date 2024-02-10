CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
  id              UUID         PRIMARY KEY DEFAULT uuid_generate_v4(),
  email           VARCHAR      UNIQUE NOT NULL,
  password_hash   VARCHAR      NOT NULL
);

CREATE TABLE IF NOT EXISTS questions (
  id              BIGSERIAL   PRIMARY KEY,
  user_id         UUID        NOT NULL,
  text            VARCHAR     NOT NULL,
  created_at      TIMESTAMP   DEFAULT now(),

  CONSTRAINT fk_questions_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);
