-- Modify "transaction" table
ALTER TABLE "transaction" ADD COLUMN "creator_id" uuid NULL, ADD COLUMN "updator_id" uuid NULL, ADD
 CONSTRAINT "fk_transaction_creator" FOREIGN KEY ("creator_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION, ADD
 CONSTRAINT "fk_transaction_last_updated_user" FOREIGN KEY ("updator_id") REFERENCES "user" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
