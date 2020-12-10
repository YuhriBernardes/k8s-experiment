package main

import (
	"mykube/database"
	"mykube/server"
	"mykube/utils"

	log "github.com/sirupsen/logrus"
)

func configLogs() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:          false,
		DisableQuote:           false,
		DisableLevelTruncation: false,
	})
	log.SetLevel(log.InfoLevel)
}

func main() {
	configLogs()
	dbConfig := database.DbConfig{
		Port:     utils.GetEnvUint("DB_PORT"),
		Host:     utils.GetEnvString("DB_HOST"),
		DbName:   utils.GetEnvString("DB_NAME"),
		User:     utils.GetEnvString("DB_USER"),
		Password: utils.GetEnvString("DB_PASS"),
	}
	log.WithFields(log.Fields{
		"host": dbConfig.Host,
		"port": dbConfig.Port,
		"name": dbConfig.DbName,
	}).Info("Starting db")
	db, err := database.NewDb(dbConfig)

	if err := db.Migrate(); err != nil {
		log.WithError(err).Error("Failed to execute migrations")
	}

	if err != nil {
		log.WithError(err).Error("Failed to start database")
	}

	srvr := server.NewServer(db)
	srvr.CreateUser()
	srvr.GetAll()

	srvr.Start()
}
