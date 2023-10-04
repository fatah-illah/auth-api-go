-- name: CreateUser :one
INSERT INTO
    users(
        username,
        password_hash,
        email
    )
VALUES ($1, $2, $3) RETURNING id;

-- name: GetUserByUsername :one
SELECT id, password_hash FROM users WHERE username = $1;

-- name: CreateTokenForUser :one
INSERT INTO
    tokens(
        user_id,
        access_token,
        expires_at
    )
VALUES ($1, $2, $3) RETURNING id;

-- name: VerifyAccessToken :one
SELECT user_id FROM tokens WHERE access_token = $1 AND expires_at > CURRENT_TIMESTAMP;

-- name: GetProtectedResource :one
SELECT * FROM protected_resources WHERE resource_id = $1;
