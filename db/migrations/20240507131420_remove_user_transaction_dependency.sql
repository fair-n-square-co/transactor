-- Modify "user" table
ALTER TABLE "user" DROP CONSTRAINT "fk_transaction_creator", DROP CONSTRAINT "fk_transaction_last_updated_user";
