
CREATE TABLE IF NOT EXISTS telegram_users (
  user_id         UUID        NOT NULL,
  chat_id         BIGINT      PRIMARY KEY
);

ALTER TABLE questions DROP CONSTRAINT fk_questions_user_id;

