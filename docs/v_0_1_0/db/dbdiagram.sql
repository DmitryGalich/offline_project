CREATE TABLE "user" (
  "id" uuid PRIMARY KEY,
  "login" varchar,
  "password" varchar,
  "mail" varchar,
  "phone" varchar,
  "first_name" varchar,
  "second_name" varchar,
  "date_of_birth" date,
  "created_at" timestamp
);

CREATE TABLE "chat" (
  "id" uuid PRIMARY KEY
);

CREATE TABLE "chat_comment" (
  "id" uuid PRIMARY KEY,
  "chat" uuid,
  "user" uuid
);

ALTER TABLE "chat_comment" ADD FOREIGN KEY ("chat") REFERENCES "chat" ("id");

ALTER TABLE "chat_comment" ADD FOREIGN KEY ("user") REFERENCES "user" ("id");
