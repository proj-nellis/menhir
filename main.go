package main

import (
	"log"

	"projnellis.com/menhir/app"
	"projnellis.com/menhir/db"
)

func main() {
	app := app.Init()
	newAccount, err := app.Database.Queries.CreateAccount(app.Database.Context, db.CreateAccountParams{
		ID:       "232323",
		Email:    "guamjust@gmail.com",
		Password: "acASCascaScasc",
	})
	if err != nil {
		return
	}
	log.Println(newAccount)
}
