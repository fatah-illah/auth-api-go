// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package models

import (
	"context"
	"database/sql"
)

const createTokenForUser = `-- name: CreateTokenForUser :one
INSERT INTO
    tokens(
        user_id,
        access_token,
        expires_at
    )
VALUES ($1, $2, $3) RETURNING id
`

type CreateTokenForUserParams struct {
	UserID      sql.NullInt32 `db:"user_id" json:"userId"`
	AccessToken string        `db:"access_token" json:"accessToken"`
	ExpiresAt   sql.NullTime  `db:"expires_at" json:"expiresAt"`
}

func (q *Queries) CreateTokenForUser(ctx context.Context, arg CreateTokenForUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createTokenForUser, arg.UserID, arg.AccessToken, arg.ExpiresAt)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO
    users(
        username,
        password_hash,
        email
    )
VALUES ($1, $2, $3) RETURNING id
`

type CreateUserParams struct {
	Username     string `db:"username" json:"username"`
	PasswordHash string `db:"password_hash" json:"passwordHash"`
	Email        string `db:"email" json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.PasswordHash, arg.Email)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getProtectedResource = `-- name: GetProtectedResource :one
SELECT resource_id, resource_name, resource_data FROM protected_resources WHERE resource_id = $1
`

func (q *Queries) GetProtectedResource(ctx context.Context, resourceID int32) (ProtectedResource, error) {
	row := q.db.QueryRowContext(ctx, getProtectedResource, resourceID)
	var i ProtectedResource
	err := row.Scan(&i.ResourceID, &i.ResourceName, &i.ResourceData)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, password_hash FROM users WHERE username = $1
`

type GetUserByUsernameRow struct {
	ID           int32  `db:"id" json:"id"`
	PasswordHash string `db:"password_hash" json:"passwordHash"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(&i.ID, &i.PasswordHash)
	return i, err
}

const verifyAccessToken = `-- name: VerifyAccessToken :one
SELECT user_id FROM tokens WHERE access_token = $1 AND expires_at > CURRENT_TIMESTAMP
`

func (q *Queries) VerifyAccessToken(ctx context.Context, accessToken string) (sql.NullInt32, error) {
	row := q.db.QueryRowContext(ctx, verifyAccessToken, accessToken)
	var user_id sql.NullInt32
	err := row.Scan(&user_id)
	return user_id, err
}
