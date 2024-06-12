
INSERT INTO "store_type" ("id", "type")
VALUES
  ((SELECT id FROM "store" WHERE "name" = 'The Clucking Bell'), 'RESTAURANT'),
  ((SELECT id FROM "store" WHERE "name" = 'Up-n-Atom Burger'), 'RESTAURANT'),
  ((SELECT id FROM "store" WHERE "name" = '24/7'), 'MARKET'),
  ((SELECT id FROM "store" WHERE "name" = 'Pill Pharm'), 'PHARMACY'),
  ((SELECT id FROM "store" WHERE "name" = 'Rusty Brown''s Ring Donuts'), 'TOBBACO'),
  ((SELECT id FROM "store" WHERE "name" = 'Liquor Ace'), 'MARKET'),
  ((SELECT id FROM "store" WHERE "name" = 'Bean Machine'), 'RESTAURANT'),
  ((SELECT id FROM "store" WHERE "name" = 'Yellow Jack Inn'), 'PUB'),
  ((SELECT id FROM "store" WHERE "name" = 'Discount Store'), 'MARKET'),
  ((SELECT id FROM "store" WHERE "name" = '24/7 Supermarket'), 'CONVENIENCE');
