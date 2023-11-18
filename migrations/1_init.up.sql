CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
  "login" varchar NOT NULL
  -- "password" varchar NOT NULL,
  -- "mail" varchar NOT NULL,
  -- "phone" varchar NOT NULL,
  -- "first_name" varchar NOT NULL,
  -- "second_name" varchar NOT NULL,
  -- "date_of_birth" date NOT NULL,
  -- "created_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE "chats" (
   "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY
);

-- CREATE TABLE "chat_comments" (
--   "id" UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
--   "chat" UUID DEFAULT uuid_generate_v4() NOT NULL,
--   "author" UUID DEFAULT uuid_generate_v4() NOT NULL,
--   "created_at" timestamp NOT NULL DEFAULT NOW()
-- );

-- ALTER TABLE "chat_comments" ADD FOREIGN KEY ("chat") REFERENCES "chats" ("id");

-- ALTER TABLE "chat_comments" ADD FOREIGN KEY ("author") REFERENCES "users" ("id");
