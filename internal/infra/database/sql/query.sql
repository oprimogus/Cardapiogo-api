-- name: CreateUser :exec
INSERT INTO cardapiogo.users (email, password, role, account_provider, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: CreateUserWithOAuth :exec
INSERT INTO cardapiogo.users (email, role, account_provider, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW());

-- name: GetUserById :one
SELECT id, profile_id, email, password, role, account_provider, created_at, updated_at FROM cardapiogo.users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, profile_id, email, password, role, account_provider, created_at, updated_at FROM cardapiogo.users
WHERE email = $1
LIMIT 1;

-- name: GetUser :many
SELECT id, profile_id, email, role, account_provider, created_at, updated_at FROM cardapiogo.users
ORDER BY created_at desc
LIMIT $1 OFFSET $2;

-- name: UpdateUser :exec
UPDATE cardapiogo.users
SET 
    email = $2,
    role = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE cardapiogo.users
SET 
    password = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserProfile :exec
UPDATE cardapiogo.users
SET 
    profile_id = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: CreateProfileAndReturnID :one
INSERT INTO cardapiogo.profile (name, last_name, cpf, phone, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING id;

-- name: GetProfileByID :one
SELECT id, name, last_name, cpf, phone FROM cardapiogo.profile
WHERE id = $1
LIMIT 1;

-- name: GetProfileByUserID :one
SELECT id, name, last_name, cpf, phone FROM cardapiogo.profile p
inner join (SELECT profile_id from cardapiogo.users where users.id = $1) u on p.id = u.profile_id
LIMIT 1;

-- name: UpdateProfile :exec
UPDATE cardapiogo.profile
SET 
    name = $2,
    last_name = $3,
    phone = $4,
    updated_at = NOW()
WHERE id = (
    SELECT profile_id
    FROM users
    WHERE users.id = $1
);

-- name: UpdateProfileCpf :exec
UPDATE cardapiogo.profile
SET 
    cpf = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: CreateAddressOfProfile :exec
INSERT INTO cardapiogo.address (profile_id, street, number, complement, district, zip_code, city, state, latitude, longitude, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
RETURNING ID;

-- name: CreateAddressOfStore :exec
INSERT INTO cardapiogo.address (street, number, complement, district, zip_code, city, state, latitude, longitude, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
RETURNING ID;

-- name: LinkAddressInStore :exec
UPDATE cardapiogo.store 
SET 
    address_id = $1,
    updated_at = NOW()
WHERE id = $2;

--name GetAddressByProfile :many
SELECT street, number, complement, district, zip_code, city, state, latitude, longitude FROM cardapiogo.address
WHERE profile_id = $1;

--name GetAddressById :one
SELECT street, number, complement, district, zip_code, city, state, latitude, longitude FROM cardapiogo.address
WHERE id = $1
LIMIT 1;

-- name: GetStoresListByFilter :many
SELECT s.id, s.name, s.score, s.district
FROM cardapiogo.store s
INNER JOIN cardapiogo.store_restaurant_type srt on srt.store_id = s.id
WHERE
    ($3 is NULL or s.name like '%' || $3 || '%')
    AND ($4 is NULL OR s.type = $4)
    AND ($5 is NULL OR srt.restaurant_type = ANY($5::string[]))
ORDER BY s.score
LIMIT $1 OFFSET $2;

-- name: CreateStore :one
INSERT INTO cardapiogo.store (name, cpf_cnpj, phone, type, latitude, longitude, street, number, complement, district, zip_code, city, state, created_at, updated_at) 
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW(), NOW())
RETURNING id;



-- name: AddOwner :exec
INSERT INTO cardapiogo.owner (profile_id, store_id, created_at)
VALUES ($1, $2, NOW());


