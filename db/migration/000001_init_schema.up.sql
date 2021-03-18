CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "cpf" varchar NOT NULL,
  "secret" varchar NOT NULL,
  "balance" bigint,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "account_origin_id" bigint NOT NULL,
  "account_destination_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "transfers" ADD FOREIGN KEY ("account_origin_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("account_destination_id") REFERENCES "accounts" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

CREATE INDEX ON "accounts" ("name");

CREATE INDEX ON "accounts" ("id");

CREATE INDEX ON "transfers" ("account_origin_id");

CREATE INDEX ON "transfers" ("account_destination_id");

CREATE INDEX ON "transfers" ("account_origin_id", "account_destination_id");

CREATE INDEX ON "entries" ("account_id");