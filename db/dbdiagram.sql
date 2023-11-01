CREATE TABLE "user" (
  "id" uuid PRIMARY KEY uuid_generate_v4(),
  "login" varchar NOT NULL,
  "password" varchar NOT NULL,
  "mail" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "date_of_birth" date NOT NULL,
  "created_at" timestamp NOT NULL now()
);

CREATE TABLE "chat" (
  "id" uuid PRIMARY KEY uuid_generate_v4()
);

CREATE TABLE "chat_comment" (
  "id" uuid PRIMARY KEY uuid_generate_v4(),
  "chat" uuid NOT NULL,
  "author" uuid NOT NULL
);

ALTER TABLE "chat_comment" ADD FOREIGN KEY ("chat") REFERENCES "chat" ("id");

ALTER TABLE "chat_comment" ADD FOREIGN KEY ("author") REFERENCES "user" ("id");
