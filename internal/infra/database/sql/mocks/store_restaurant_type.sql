INSERT INTO "cardapiogo"."store_restaurant_type" ("id", "restaurant_type")
VALUES
  ((SELECT id FROM "cardapiogo"."store" WHERE "name" = 'The Clucking Bell'), 'AMERICAN'),
  ((SELECT id FROM "cardapiogo"."store" WHERE "name" = 'Up-n-Atom Burger'), 'AMERICAN'),
  ((SELECT id FROM "cardapiogo"."store" WHERE "name" = 'Bean Machine'), 'ITALIAN'),
  ((SELECT id FROM "cardapiogo"."store" WHERE "name" = 'The Clucking Bell'), 'MEXICAN'),
  ((SELECT id FROM "cardapiogo"."store" WHERE "name" = 'Bean Machine'), 'AMERICAN');

