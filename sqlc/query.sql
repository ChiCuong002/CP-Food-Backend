-- name: GetUserByEmail :one
SELECT email FROM users WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserObjectByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpsertRefreshToken :one
INSERT INTO keys (user_id, refresh_token) VALUES ($1, $2) 
ON CONFLICT (user_id) 
DO UPDATE 
SET refresh_token = EXCLUDED.refresh_token, 
used_refresh_token = ARRAY_APPEND(keys.used_refresh_token, keys.refresh_token)
RETURNING *;   

-- name: RemoveRefreshToken :exec
DELETE FROM keys 
WHERE user_id = $1;

-- name: GetUserTokenById :one
SELECT * 
FROM users us, keys k
WHERE us.id = k.user_id 
AND us.id = $1;