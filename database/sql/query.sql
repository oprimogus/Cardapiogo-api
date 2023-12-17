-- name: CreateUser :exec
INSERT INTO cardapio.users (email, password, role, account_provider, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: GetUser :one
SELECT * FROM cardapio.users
WHERE id = $1
LIMIT 1;

-- name: UpdateUser :exec
UPDATE cardapio.users
SET 
    email = $2,
    role = $3,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE cardapio.users
SET 
    password = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateUserProfile :exec
UPDATE cardapio.users
SET 
    profile_id = $2,
    updated_at = NOW()
WHERE id = $1;

-- name: CreateProfile :exec
INSERT INTO cardapio.profile (name, last_name, cpf, phone, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: GetProfile :one
SELECT * FROM cardapio.profile
WHERE id = $1
LIMIT 1;

-- name: UpdateProfile :exec
UPDATE cardapio.profile
SET 
    name = $2,
    last_name = $3,
    phone = $4,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateProfileCpf :exec
UPDATE cardapio.profile
SET 
    cpf = $2,
    updated_at = NOW()
WHERE id = $1;



