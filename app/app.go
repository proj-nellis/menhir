package app

import (
	"log"
	"os"

	"github.com/bwmarrin/snowflake"
	menhir_config "projnellis.com/menhir/config"
)

type App struct {
	Database   *Database
	Config     menhir_config.Config
	Snowflakes *snowflake.Node
	Mailer     *Mailer
	Handlebars *Handlebars
}

func Init() App {
	config, err := menhir_config.Init()
	if err != nil {
		log.Fatal("Failed to initialize the config :(")
		os.Exit(-1)
	}

	node, err := snowflake.NewNode(config.NodeId)
	if err != nil {
		log.Fatal("Failed to initialize snowflake id generator")
		os.Exit(-1)
	}

	// Don't ask why :()
	RunMigrate(config.Database)

	database := &Database{}
	database.Init(config.Database)

	mailer := &Mailer{}
	mailer.Init(config.Mg.Domain, config.Mg.Pkey)

	handlebars := &Handlebars{}
	handlebars.Init("./templates")

	return App{
		Database:   database,
		Config:     config,
		Snowflakes: node,
		Handlebars: handlebars,
		Mailer:     mailer,
	}
}
