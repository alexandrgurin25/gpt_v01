
CREATE TABLE IF NOT EXISTS telegram_users (
  user_id         UUID        NOT NULL,
  chat_id         BIGINT      PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS questions (
  id              BIGSERIAL   PRIMARY KEY,
  user_id         UUID        NOT NULL,
  text            VARCHAR     NOT NULL,
  created_at      TIMESTAMP   DEFAULT now(),

  CONSTRAINT fk_questions_user_id FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);

ALTER TABLE questions DROP CONSTRAINT fk_questions_user_id;