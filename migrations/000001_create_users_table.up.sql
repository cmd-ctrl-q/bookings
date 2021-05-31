BEGIN;

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "first_name" varchar(255) NOT NULL DEFAULT '',
  "last_name" varchar(255) NOT NULL DEFAULT '',
  "email" varchar(255) NOT NULL,
  "password" varchar(60),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "access_level" integer NOT NULL DEFAULT 1
);

COMMIT;