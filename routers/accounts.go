package routers

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

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

	new_account, err := r_ctx.App.Database.Queries.CreateAccount(r_ctx.App.Database.Context, db.CreateAccountParams{
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

	message := (*r_ctx.App.Mailer).Mg.NewMessage(r_ctx.App.Config.Mg.From, "Account creation: Email verification", "", new_account.Email)
	message.SetHtml(r_ctx.App.Handlebars.Render("email_verification", map[string]any{
		"link": fmt.Sprintf("http://localhost:8080/accounts/email/verify?tok=%s", new_account.EmailVerificationToken.String),
	}))

	email_ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, id, err := (*r_ctx.App.Mailer).Mg.Send(email_ctx, message)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, struct{ Message string }{Message: "Failed to send email..."})
		/// TODO: Probably delete/invalidate previously created request
		return
	}
	fmt.Printf("SENT EMAIL: [ID: %s Resp: %s]\n", id, resp)
	ctx.Status(http.StatusCreated)
}
