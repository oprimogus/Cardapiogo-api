-- name: CreateUser :exec
INSERT INTO users (email, password, role, account_provider, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: CreateUserWithOAuth :exec
INSERT INTO users (email, role, account_provider, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW());

-- name: GetUserById :one
SELECT id, profile_id, email, password, role, account_provider, created_at, updated_at FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, profile_id, email, password, role, account_provider, created_at, updated_at FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUser :many
SELECT id, profile_id, email, role, account_provider, created_at, updated_at FROM users
ORDER BY created_at desc
LIMIT $1 OFFSET $2;

-- name: UpdateUser :exec
UPDATE users
SET 
    email = $2,
    role = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET 
    password = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserProfile :exec
UPDATE users
SET 
    profile_id = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: CreateProfileAndReturnID :one
INSERT INTO profile (name, last_name, cpf, phone, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING id;

-- name: GetProfileByID :one
SELECT id, name, last_name, cpf, phone FROM profile
WHERE id = $1
LIMIT 1;

-- name: GetProfileByUserID :one
SELECT id, name, last_name, cpf, phone FROM profile p
inner join (SELECT profile_id from users u where u.id = $1) u on p.id = u.profile_id
LIMIT 1;

-- name: UpdateProfile :exec
UPDATE profile
SET 
    name = $2,
    last_name = $3,
    phone = $4,
    updated_at = NOW()
WHERE id = (
    SELECT profile_id
    FROM users u
    WHERE u.id = $1
);

-- name: UpdateProfileCpf :exec
UPDATE profile
SET 
    cpf = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: CreateAddressOfProfile :exec
INSERT INTO address (profile_id, street, number, complement, district, zip_code, city, state, latitude, longitude, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, NOW(), NOW())
RETURNING ID;

-- name: CreateAddressOfStore :exec
INSERT INTO address (street, number, complement, district, zip_code, city, state, latitude, longitude, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(), NOW())
RETURNING ID;

-- name: LinkAddressInStore :exec
UPDATE store 
SET 
    address_id = $1,
    updated_at = NOW()
WHERE id = $2;


