CREATE TABLE "posts" (
  "id" uuid PRIMARY KEY,
  "post" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "created_by" varchar NOT NULL DEFAULT 'SYSTEM',
  "updated_at" timestamptz,
  "updated_by" varchar
);

CREATE INDEX ON "posts" ("id");
