CREATE TABLE "users" (
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

CREATE TABLE "chats" (
  "id" uuid PRIMARY KEY uuid_generate_v4()
);

CREATE TABLE "chat_comments" (
  "id" uuid PRIMARY KEY uuid_generate_v4(),
  "chat" uuid NOT NULL,
  "author" uuid NOT NULL,
  "created_at" timestamp NOT NULL now()
);

ALTER TABLE "chat_comments" ADD FOREIGN KEY ("chats") REFERENCES "chats" ("id");

ALTER TABLE "chat_comments" ADD FOREIGN KEY ("authors") REFERENCES "users" ("id");
