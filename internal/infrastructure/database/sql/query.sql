-- name: CreateStore :exec
INSERT INTO store (id, cpf_cnpj, owner_id, name, active, phone, score, type, address_line_1, address_line_2, neighborhood, city, state, postal_code,
  latitude, longitude, country, created_at, updated_at)
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, NOW(), NOW());
