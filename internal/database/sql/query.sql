-- name: CreateStore :exec
INSERT INTO store (id, cpf_cnpj, owner_id, name, active, phone, score, type, address_line_1, address_line_2, neighborhood, city, state, postal_code,
  latitude, longitude, country, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, NOW() AT TIME ZONE 'UTC', NOW() AT TIME ZONE 'UTC');

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
    updated_at = NOW() AT TIME ZONE 'UTC'
WHERE id = $1 AND owner_id = $2;

-- name: SetProfileImage :exec
UPDATE store
  SET 
    profile_image = $2
WHERE id = $1;

-- name: SetHeaderImage :exec
UPDATE store
  SET 
    header_image = $2
WHERE id = $1;

-- name: IsOwner :one
SELECT EXISTS(SELECT 1 FROM store WHERE id = $1 AND owner_id = $2);

-- name: GetStoreByID :one
SELECT s.id, s.name, s.phone, s.score, s.type, s.address_line_1, 
s.address_line_2, s.neighborhood, s.city, s.state, s.country, s.profile_image, s.header_image
FROM store s
WHERE id = $1;

-- name: GetStoreBusinessHoursByID :many
SELECT week_day, timezone, opening_time, closing_time
FROM business_hour
WHERE store_id = $1
ORDER BY week_day;

-- name: GetStoreByFilter :many
SELECT s.id, s.name, s.score, s.type, s.neighborhood, s.latitude, s.longitude, s.profile_image
FROM store s
WHERE 1 = 1
  AND (COALESCE(NULLIF($1, ''), s.name) IS NULL OR s.name LIKE '%' || COALESCE(NULLIF($1, ''), s.name) || '%')
  AND (COALESCE($2, s.score) IS NULL OR s.score >= COALESCE($2, s.score))
  AND (COALESCE(NULLIF($3, '')::"ShopType", s.type) IS NULL OR s.type = COALESCE(NULLIF($3, '')::"ShopType", s.type))
  AND (COALESCE(NULLIF($4, ''), s.city) IS NULL OR s.city = COALESCE(NULLIF($4, ''), s.city))
ORDER BY s.score DESC, s.type;

-- name: AddBusinessHours :batchexec
INSERT INTO business_hour(store_id, week_day, opening_time, closing_time, timezone)
VALUES ($1, $2, $3, $4, $5);

-- name: DeleteBusinessHours :batchexec
DELETE FROM business_hour
WHERE store_id = $1
  AND week_day = $2
  AND opening_time = $3
  AND closing_time = $4;

-- name: FindStoreBusinessHoursByStoreId :many
SELECT bh.store_id, bh.week_day, bh.opening_time, bh.closing_time, bh.timezone
FROM business_hour bh
WHERE 1 = 1
  AND bh.store_id = ANY($1::UUID[]);