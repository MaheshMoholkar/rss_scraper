-- name: CreateUser :one
INSERT INTO users (created_at, updated_at, name, api_key)
VALUES($1, $2, $3,
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserByAPIKEY :one

SELECT * FROM users WHERE api_key = $1;