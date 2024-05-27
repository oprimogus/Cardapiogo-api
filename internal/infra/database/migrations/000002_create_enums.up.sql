CREATE TYPE "cardapiogo"."user_role" AS ENUM (
  'CUSTOMER',
  'OWNER',
  'EMPLOYEE',
  'DELIVERY_MAN',
  'ADMIN'
);

CREATE TYPE "cardapiogo"."cousine_type" AS ENUM (
  'ITALIAN',
  'JAPANESE',
  'MEXICAN',
  'ARABIC',
  'BRAZILIAN',
  'THAI',
  'AMERICAN'
);

CREATE TYPE "cardapiogo"."shop_type" AS ENUM (
  'RESTAURANT',
  'PHARMACY',
  'TOBBACO',
  'MARKET',
  'CONVENIENCE',
  'PUB'
);

CREATE TYPE "cardapiogo"."payment_form" AS ENUM (
  'CREDIT_CARD',
  'DEBIT_CARD',
  'PIX',
  'CASH'
);

CREATE TYPE "cardapiogo"."order_status" AS ENUM (
  'CREATED',
  'ACCEPTED',
  'IN_PROGRESS',
  'FINISHED',
  'CANCELED'
);

CREATE TYPE "cardapiogo"."weekday" AS ENUM (
  'MONDAY',
  'TUESDAY',
  'WEDNESDAY',
  'THURSDAY',
  'FRIDAY',
  'SATURDAY',
  'SUNDAY'
);

CREATE TYPE "cardapiogo"."account_provider" AS ENUM (
  'GOOGLE',
  'APPLE',
  'META',
  'CARDAPIOGO'
);
