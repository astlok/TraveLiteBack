package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"net/http"
	"travalite/configs"
	"travalite/internal/app"
	"travalite/pkg/logger"
	"travalite/pkg/postgresql"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/toml/server.toml", "path to config file")
}

func main() {
	flag.Parse()
	config := configs.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	err = logger.InitLogger("stdout", config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := postgresql.NewPostgres(config.DSN)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := postgres.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	router := app.ConfigureRoute(*config, postgres.GetPostgres())

	server := http.Server{
		Addr:    config.BindAddr,
		Handler: router,
	}

	if config.HTTPS {
		log.Println("TLS server starting at port: ", server.Addr)
		if err := server.ListenAndServeTLS(
			"/etc/letsencrypt/live/findfreelancer.ru/cert.pem",
			"/etc/letsencrypt/live/findfreelancer.ru/privkey.pem"); err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Server starting at port", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
