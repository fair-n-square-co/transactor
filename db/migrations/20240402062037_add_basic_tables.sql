-- Create "transaction_user" table
CREATE TABLE "transaction_user" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "transaction_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "amount_currency_code" text NULL,
  "amount_units" bigint NULL,
  "amount_nanos" integer NULL,
  "transaction_user_type" "transaction_user_type" NOT NULL DEFAULT 'payer',
  PRIMARY KEY ("id")
);
-- Create index "idx_transaction_user_deleted_at" to table: "transaction_user"
CREATE INDEX "idx_transaction_user_deleted_at" ON "transaction_user" ("deleted_at");
-- Create "transaction" table
CREATE TABLE "transaction" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "amount_currency_code" text NULL,
  "amount_units" bigint NULL,
  "amount_nanos" integer NULL,
  "description" text NOT NULL,
  "type" "transaction_type" NOT NULL DEFAULT 'payment',
  PRIMARY KEY ("id")
);
-- Create index "idx_transaction_deleted_at" to table: "transaction"
CREATE INDEX "idx_transaction_deleted_at" ON "transaction" ("deleted_at");
-- Create "user" table
CREATE TABLE "user" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "email" text NOT NULL,
  "first_name" text NOT NULL,
  "last_name" text NULL,
  "phone" character varying(100) NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_transaction_creator" FOREIGN KEY ("id") REFERENCES "transaction" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_transaction_last_updated_user" FOREIGN KEY ("id") REFERENCES "transaction" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_user_deleted_at" to table: "user"
CREATE INDEX "idx_user_deleted_at" ON "user" ("deleted_at");
-- Create index "uni_user_email" to table: "user"
CREATE UNIQUE INDEX "uni_user_email" ON "user" ("email");
-- Create index "uni_user_phone" to table: "user"
CREATE UNIQUE INDEX "uni_user_phone" ON "user" ("phone");
-- Create "group" table
CREATE TABLE "group" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_group_deleted_at" to table: "group"
CREATE INDEX "idx_group_deleted_at" ON "group" ("deleted_at");
-- Create "user_groups" table
CREATE TABLE "user_groups" (
  "group_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "user_id" uuid NOT NULL DEFAULT gen_random_uuid(),
  PRIMARY KEY ("group_id", "user_id"),
  CONSTRAINT "fk_user_groups_group" FOREIGN KEY ("group_id") REFERENCES "group" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_user_groups_user" FOREIGN KEY ("user_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
