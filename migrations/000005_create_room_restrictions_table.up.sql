BEGIN;

CREATE TABLE "room_restrictions" (
  "id" serial PRIMARY KEY,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "room_id" integer NOT NULL,
  "reservation_id" integer,
  "restriction_id" integer
);

COMMIT;