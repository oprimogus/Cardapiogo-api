CREATE TYPE "user_role" AS ENUM (
  'CUSTOMER',
  'OWNER',
  'EMPLOYEE',
  'DELIVERY_MAN',
  'ADMIN'
);

CREATE TYPE "account_provider" AS ENUM (
  'GOOGLE',
  'APPLE',
  'META',
  'CARDAPIOGO'
);

CREATE TYPE "cousine_type" AS ENUM (
  'ITALIAN',
  'JAPANESE',
  'MEXICAN',
  'ARABIC',
  'BRAZILIAN',
  'THAI',
  'AMERICAN'
);

CREATE TYPE "shop_type" AS ENUM (
  'RESTAURANT',
  'PHARMACY',
  'TOBBACO',
  'MARKET',
  'CONVENIENCE',
  'PUB'
);

CREATE TYPE "payment_form" AS ENUM (
  'CREDIT_CARD',
  'DEBIT_CARD',
  'PIX',
  'CASH'
);

CREATE TYPE "order_status" AS ENUM (
  'CREATED',
  'ACCEPTED',
  'IN_PROGRESS',
  'FINISHED',
  'CANCELED'
);

CREATE TYPE "weekday" AS ENUM (
  'MONDAY',
  'TUESDAY',
  'WEDNESDAY',
  'THURSDAY',
  'FRIDAY',
  'SATURDAY',
  'SUNDAY'
);