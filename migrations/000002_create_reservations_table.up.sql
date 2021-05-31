BEGIN;

CREATE TABLE "reservations" (
  "id" integer PRIMARY KEY,
  "first_name" varchar(255) NOT NULL DEFAULT '',
  "last_name" varchar(255) NOT NULL DEFAULT '',
  "email" varchar(255) NOT NULL,
  "phone" varchar(255) NOT NULL DEFAULT '',
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "room_id" integer NOT NULL
);

COMMIT;