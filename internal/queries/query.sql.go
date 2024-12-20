// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package queries

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email, password, status, role_id, created_at
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.RoleID,
		&i.CreatedAt,
	)
	return i, err
}

const detailUser = `-- name: DetailUser :one
SELECT users.id, users.name, email, roles.name as role
FROM users, roles
WHERE users.role_id = roles.id
AND users.id = $1
`

type DetailUserRow struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (q *Queries) DetailUser(ctx context.Context, id int32) (DetailUserRow, error) {
	row := q.db.QueryRowContext(ctx, detailUser, id)
	var i DetailUserRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Role,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT email FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	err := row.Scan(&email)
	return email, err
}

const getUserObjectByEmail = `-- name: GetUserObjectByEmail :one
SELECT id, name, email, password, status, role_id, created_at FROM users WHERE email = $1
`

func (q *Queries) GetUserObjectByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserObjectByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.RoleID,
		&i.CreatedAt,
	)
	return i, err
}

const getUserTokenById = `-- name: GetUserTokenById :one
SELECT us.id, name, email, password, status, role_id, created_at, k.id, user_id, refresh_token, used_refresh_token 
FROM users us, keys k
WHERE us.id = k.user_id 
AND us.id = $1
`

type GetUserTokenByIdRow struct {
	ID               int32          `json:"id"`
	Name             string         `json:"name"`
	Email            string         `json:"email"`
	Password         string         `json:"password"`
	Status           NullUserStatus `json:"status"`
	RoleID           sql.NullInt32  `json:"role_id"`
	CreatedAt        sql.NullTime   `json:"created_at"`
	ID_2             int32          `json:"id_2"`
	UserID           int32          `json:"user_id"`
	RefreshToken     sql.NullString `json:"refresh_token"`
	UsedRefreshToken []string       `json:"used_refresh_token"`
}

func (q *Queries) GetUserTokenById(ctx context.Context, id int32) (GetUserTokenByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getUserTokenById, id)
	var i GetUserTokenByIdRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Status,
		&i.RoleID,
		&i.CreatedAt,
		&i.ID_2,
		&i.UserID,
		&i.RefreshToken,
		pq.Array(&i.UsedRefreshToken),
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
WITH total_count AS (
  SELECT COUNT(*) AS count
  FROM users
  WHERE 
    ($3::text IS NULL OR name ILIKE '%' || $3 || '%') AND
    ($4::text IS NULL OR status::text ILIKE '%' || $4 || '%')
)
SELECT 
  id, name, email, status, created_at, total_count.count AS total_count
FROM 
  users, total_count
WHERE 
  ($3::text IS NULL OR name ILIKE '%' || $3 || '%') AND
  ($4::text IS NULL OR status::text ILIKE '%' || $4 || '%')
ORDER BY
  CASE WHEN $5 = 'iddesc' THEN id END DESC,
  CASE WHEN $5 = 'idasc' THEN id END ASC
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
	Search string      `json:"search"`
	Status string      `json:"status"`
	Sort   interface{} `json:"sort"`
}

type ListUsersRow struct {
	ID         int32          `json:"id"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Status     NullUserStatus `json:"status"`
	CreatedAt  sql.NullTime   `json:"created_at"`
	TotalCount int64          `json:"total_count"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, listUsers,
		arg.Limit,
		arg.Offset,
		arg.Search,
		arg.Status,
		arg.Sort,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersRow
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Status,
			&i.CreatedAt,
			&i.TotalCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeRefreshToken = `-- name: RemoveRefreshToken :exec
DELETE FROM keys 
WHERE user_id = $1
`

func (q *Queries) RemoveRefreshToken(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, removeRefreshToken, userID)
	return err
}

const upsertRefreshToken = `-- name: UpsertRefreshToken :one
INSERT INTO keys (user_id, refresh_token) VALUES ($1, $2) 
ON CONFLICT (user_id) 
DO UPDATE 
SET refresh_token = EXCLUDED.refresh_token, 
used_refresh_token = ARRAY_APPEND(keys.used_refresh_token, keys.refresh_token)
RETURNING id, user_id, refresh_token, used_refresh_token
`

type UpsertRefreshTokenParams struct {
	UserID       int32          `json:"user_id"`
	RefreshToken sql.NullString `json:"refresh_token"`
}

func (q *Queries) UpsertRefreshToken(ctx context.Context, arg UpsertRefreshTokenParams) (Key, error) {
	row := q.db.QueryRowContext(ctx, upsertRefreshToken, arg.UserID, arg.RefreshToken)
	var i Key
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RefreshToken,
		pq.Array(&i.UsedRefreshToken),
	)
	return i, err
}
