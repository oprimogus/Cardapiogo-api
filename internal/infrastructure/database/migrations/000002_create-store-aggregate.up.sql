-- Migration 000002

CREATE TYPE "ShopType" AS ENUM (
  'restaurant',
  'pharmacy',
  'tobbaco',
  'market',
  'convenience',
  'pub'
);

CREATE TYPE "PaymentMethodEnum" AS ENUM (
  'credit',
  'debit',
  'pix',
  'cash',
  'bitcoin'
);

CREATE TABLE "store" (
  "id" uuid UNIQUE PRIMARY KEY,
  "cpf_cnpj" varchar UNIQUE NOT NULL,
  "owner_id" uuid NOT NULL,
  "name" varchar UNIQUE NOT NULL,
  "active" bool NOT NULL,
  "phone" varchar UNIQUE NOT NULL,
  "score" int NOT NULL,
  "type" "ShopType" NOT NULL,
  "address_line_1" varchar NOT NULL,
  "address_line_2" varchar NOT NULL,
  "neighborhood" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "postal_code" varchar NOT NULL,
  "latitude" varchar,
  "longitude" varchar,
  "country" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "bussiness_hour" (
  "store_id" uuid PRIMARY KEY,
  "week_day" int NOT NULL,
  "opening_time" varchar NOT NULL,
  "closing_time" varchar NOT NULL
);

CREATE TABLE "payment_method" (
  "id" integer PRIMARY KEY,
  "method" "PaymentMethodEnum"
);

CREATE TABLE "store_payment_method" (
  "store_id" uuid PRIMARY KEY,
  "payment_method_id" integer
);

CREATE TABLE "address" (
  "id" integer PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "address_line_1" varchar NOT NULL,
  "address_line_2" varchar NOT NULL,
  "neighborhood" varchar NOT NULL,
  "city" varchar NOT NULL,
  "state" varchar NOT NULL,
  "postal_code" varchar NOT NULL,
  "latitude" varchar,
  "longitude" varchar NOT NULL,
  "country" varchar NOT NULL,
  "created_at" timestamp NOT NULL
);

COMMENT ON COLUMN "store"."owner_id" IS 'user ID of store owner';

ALTER TABLE "bussiness_hour" ADD FOREIGN KEY ("store_id") REFERENCES "store" ("id");

ALTER TABLE "store_payment_method" ADD FOREIGN KEY ("store_id") REFERENCES "store" ("id");

ALTER TABLE "store_payment_method" ADD FOREIGN KEY ("payment_method_id") REFERENCES "payment_method" ("id");

