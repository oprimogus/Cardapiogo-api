INSERT INTO "store_restaurant_type" ("id", "restaurant_type")
VALUES
  ((SELECT id FROM "store" WHERE "name" = 'The Clucking Bell'), 'AMERICAN'),
  ((SELECT id FROM "store" WHERE "name" = 'Up-n-Atom Burger'), 'AMERICAN'),
  ((SELECT id FROM "store" WHERE "name" = 'Bean Machine'), 'ITALIAN'),
  ((SELECT id FROM "store" WHERE "name" = 'The Clucking Bell'), 'MEXICAN'),
  ((SELECT id FROM "store" WHERE "name" = 'Bean Machine'), 'AMERICAN');

