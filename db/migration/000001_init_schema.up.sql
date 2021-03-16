CREATE TABLE "accounts" (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "cpf" varchar NOT NULL,
  "secret" varchar NOT NULL,
  "balance" numeric(18,2),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" uuid PRIMARY KEY,
  "account_origin_id" uuid NOT NULL,
  "account_destination_id" uuid NOT NULL,
  "amount" numeric(18,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "transfers" ADD FOREIGN KEY ("account_origin_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("account_destination_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "accounts" ("id");

CREATE INDEX ON "transfers" ("account_origin_id");

CREATE INDEX ON "transfers" ("account_destination_id");

CREATE INDEX ON "transfers" ("account_origin_id", "account_destination_id");