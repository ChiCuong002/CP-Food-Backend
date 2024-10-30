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

-- name: ListUsers :many
WITH total_count AS (
  SELECT COUNT(*) AS count
  FROM users
  WHERE 
    (sqlc.arg(search)::text IS NULL OR name ILIKE '%' || sqlc.arg(search) || '%') AND
    (sqlc.arg(status)::text IS NULL OR status::text ILIKE '%' || sqlc.arg(status) || '%')
)
SELECT 
  id, name, email, status, created_at, total_count.count AS total_count
FROM 
  users, total_count
WHERE 
  (sqlc.arg(search)::text IS NULL OR name ILIKE '%' || sqlc.arg(search) || '%') AND
  (sqlc.arg(status)::text IS NULL OR status::text ILIKE '%' || sqlc.arg(status) || '%')
ORDER BY
  CASE WHEN sqlc.arg(sort) = 'iddesc' THEN id END DESC,
  CASE WHEN sqlc.arg(sort) = 'idasc' THEN id END ASC
LIMIT $1 OFFSET $2;

-- name: DetailUser :one
SELECT users.id, users.name, email, roles.name as role
FROM users, roles
WHERE users.role_id = roles.id
AND users.id = $1;