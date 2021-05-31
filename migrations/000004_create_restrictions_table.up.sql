BEGIN;

CREATE TABLE "restrictions" (
  "id" integer PRIMARY KEY,
  "restriction_name" varchar(255),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

COMMIT;