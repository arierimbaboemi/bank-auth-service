package config

import (
	"github.com/arierimbaboemi/bank-auth-service/domain"
	logger "github.com/arierimbaboemi/bank-lib-service/config"

	"github.com/jmoiron/sqlx"
)

/*
 * Use database config from .env
 */
func NewDBConnectionENV() (*sqlx.DB, error) {
	logger.GetLog().Info().Msg("Initializing database connection")
	config := &domain.Config{}
	db, err := sqlx.Connect("mysql", config.GetDatabaseENVConfig())
	if err != nil {
		logger.GetLog().Fatal().Err(err).Msg("Failed to connect database")
	} else {
		logger.GetLog().Info().Msg("Database connected")
	}

	return db, nil
}
