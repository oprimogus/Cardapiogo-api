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