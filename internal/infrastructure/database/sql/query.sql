-- name: CreateStore :exec
INSERT INTO store (id, cpf_cnpj, owner_id, name, active, phone, score, type, address_line_1, address_line_2, neighborhood, city, state, postal_code,
  latitude, longitude, country, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, NOW(), NOW());

-- name: UpdateStore :exec
UPDATE store
  SET 
    name = $3,
    phone = $4,
    type = $5,
    address_line_1 = $6,
    address_line_2 = $7,
    neighborhood = $8,
    city = $9,
    state = $10,
    postal_code = $11,
    country = $12,
    updated_at = NOW()
WHERE id = $1 AND owner_id = $2;

-- name: IsOwner :one
SELECT EXISTS(SELECT 1 FROM store WHERE id = $1 AND owner_id = $2);

-- name: GetStoreByID :one
SELECT s.id, s.name, s.phone, s.score, s.type, s.address_line_1, s.address_line_2, s.neighborhood, s.city, s.state, s.country
FROM store s
WHERE id = $1;

-- name: GetStoreBusinessHoursByID :many
SELECT week_day, opening_time, closing_time
FROM business_hour
WHERE store_id = $1
ORDER BY week_day;

-- name: UpsertBusinessHours :batchexec
INSERT INTO business_hour(store_id, week_day, opening_time, closing_time)
VALUES ($1, $2, $3, $4)
ON CONFLICT (store_id, week_day)
DO UPDATE SET
  opening_time = EXCLUDED.opening_time,
  closing_time = EXCLUDED.closing_time;

-- name: DeleteBusinessHours :batchexec
DELETE FROM business_hour
WHERE store_id = $1
  AND week_day = $2
  AND opening_time = $3
  AND closing_time = $4;