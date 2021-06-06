BEGIN;

CREATE TABLE "rooms" (
  "id" serial PRIMARY KEY,
  "room_name" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

COMMIT;