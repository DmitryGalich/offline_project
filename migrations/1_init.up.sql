CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" UUID NOT NULL DEFAULT (uuid_generate_v4()) PRIMARY KEY,
  "login" varchar NOT NULL,
  "password" varchar NOT NULL,
  "mail" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "second_name" varchar NOT NULL,
  "date_of_birth" date NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (NOW())
);

CREATE TABLE "chats" (
  "id" UUID NOT NULL DEFAULT (uuid_generate_v4()) PRIMARY KEY,
  "title" varchar NOT NULL
);

CREATE TABLE "chat_comments" (
  "id" UUID NOT NULL DEFAULT (uuid_generate_v4()) PRIMARY KEY,
  "chat_id" UUID NOT NULL,
  "author_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT NOW(),
  "edited_at" timestamp NOT NULL DEFAULT NOW(),
  "text" varchar NOT NULL
);

ALTER TABLE "chat_comments" ADD FOREIGN KEY ("chat_id") REFERENCES "chats" ("id");

ALTER TABLE "chat_comments" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");
