ALTER TABLE "bussiness_hour" DROP CONSTRAINT "bussiness_hour_store_id_fkey";
ALTER TABLE "store_payment_method" DROP CONSTRAINT "store_payment_method_store_id_fkey";
ALTER TABLE "store_payment_method" DROP CONSTRAINT "store_payment_method_payment_method_id_fkey";

DROP TABLE IF EXISTS "address" CASCADE;
DROP TABLE IF EXISTS "store_payment_method" CASCADE;
DROP TABLE IF EXISTS "payment_method" CASCADE;
DROP TABLE IF EXISTS "bussiness_hour" CASCADE;
DROP TABLE IF EXISTS "store" CASCADE;

DROP TYPE IF EXISTS "PaymentMethodEnum";
DROP TYPE IF EXISTS "ShopType";
