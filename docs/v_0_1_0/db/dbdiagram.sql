CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "login" varchar NOT NULL,
  "password" varchar NOT NULL,
  "mail" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "date_of_birth" date NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "chat" (
  "id" uuid PRIMARY KEY
);

CREATE TABLE "chat_comment" (
  "id" uuid PRIMARY KEY,
  "chat" uuid NOT NULL,
  "author" uuid NOT NULL
);

ALTER TABLE "chat_comment" ADD FOREIGN KEY ("chat") REFERENCES "chat" ("id");

ALTER TABLE "chat_comment" ADD FOREIGN KEY ("author") REFERENCES "user" ("id");
