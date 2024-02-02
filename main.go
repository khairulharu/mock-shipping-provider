package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog"
)

func main() {
	cfg := parseConfig()

	log := zerolog.New(os.Stdout)

	database, err := sql.Open("sqlite3", cfg.databasePath)
	if err != nil {
		log.Fatal().Msgf("opening sql connection: %s", err.Error())
	}
	defer func() {
		err := database.Close()
		if err != nil {
			log.Err(err).Msg("closing database connection")
		}
	}()
}
