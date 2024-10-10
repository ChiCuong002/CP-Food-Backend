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
	Name     string
	Email    string
	Password string
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
	ID               int32
	Name             string
	Email            string
	Password         string
	Status           NullUserStatus
	RoleID           sql.NullInt32
	CreatedAt        sql.NullTime
	ID_2             int32
	UserID           int32
	RefreshToken     sql.NullString
	UsedRefreshToken []string
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
	UserID       int32
	RefreshToken sql.NullString
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
