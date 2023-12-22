CREATE TABLE IF NOT EXISTS "accounts" (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    --######################################################
    "email" VARCHAR NOT NULL,
    "email_verified" BOOLEAN DEFAULT FALSE NOT NULL,
    "email_verification_token" VARCHAR UNIQUE,
    "time_email_verified" TIMESTAMP WITHOUT TIME ZONE,
    --######################################################
    "phone" VARCHAR,
    "phone_verified" BOOLEAN DEFAULT FALSE NOT NULL,
    "phone_verification_code" VARCHAR UNIQUE,
    "time_phone_verified" TIMESTAMP WITHOUT TIME ZONE,
    --######################################################
    "password" VARCHAR(10000) NOT NULL,
    "flags" BIGINT NOT NULL DEFAULT 0,
    "time_created" TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc')
);

CREATE TABLE IF NOT EXISTS "organizations" (
    "id" VARCHAR NOT NULL PRIMARY KEY,
    "owner" VARCHAR NOT NULL REFERENCES "accounts"("id"),
    "time_created" TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc')
);