package routers

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"projnellis.com/menhir/db"
)

func (r_ctx *RouteContext) PostCreateAccount(ctx *gin.Context) {
	type PostCreaetAccountData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var create_account_data PostCreaetAccountData
	if err := ctx.BindJSON(&create_account_data); err != nil {
		return
	}

	account_id := r_ctx.GenerateSnowflakeId()

	random_saltb := [16]byte{}
	rand.Read(random_saltb[:])
	hash, err := argon2id.CreateHash(create_account_data.Password, argon2id.DefaultParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, struct{ Message string }{Message: "Failed to hash password."})
		return
	}

	email_verification_tokenb := [16]byte{}
	rand.Read(email_verification_tokenb[:])
	email_verification_token := hex.EncodeToString(email_verification_tokenb[:])

	r_ctx.App.Database.Queries.CreateAccount(r_ctx.App.Database.Context, db.CreateAccountParams{
		ID:       account_id,
		Email:    create_account_data.Email,
		Password: hash,
		// IS THIS THE RIGHT WAY TO DO IT?
		EmailVerificationToken: sql.NullString{
			String: email_verification_token,
			Valid:  true,
		},
		Flags: 0,
	})

	ctx.Status(http.StatusCreated)
}
