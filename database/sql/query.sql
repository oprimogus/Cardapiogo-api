----------------------Users----------------------------
-- name: CreateUser :exec
INSERT INTO users (email, password, role, account_provider, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

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


----------------------Profile----------------------------
-- name: CreateProfile :exec
INSERT INTO profile (name, last_name, cpf, phone, created_at, updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: GetProfile :one
SELECT * FROM profile
WHERE id = $1
LIMIT 1;

-- name: UpdateProfile :exec
UPDATE profile
SET 
    name = $2,
    last_name = $3,
    phone = $4,
    updated_at = NOW()
WHERE id = $1;

-- name: UpdateProfileCpf :exec
UPDATE profile
SET 
    cpf = $2,
    updated_at = NOW()
WHERE id = $1;



