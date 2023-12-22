-- name: CreateAccount :one
INSERT INTO accounts(id, email, password) VALUES($1,$2,$3) RETURNING id;