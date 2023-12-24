-- name: CreateAccount :one
INSERT INTO accounts(id,email,password,flags,email_verification_token) VALUES($1,$2,$3,$4,$5) RETURNING id, email, flags, time_created,email_verification_token;

-- name: CreateOrganization :one
INSERT INTO organizations(id,owner) VALUES($1,$2) RETURNING id, time_created;