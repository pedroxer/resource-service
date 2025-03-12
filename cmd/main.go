package main

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/pedroxer/resource-service/internal/app"
	"github.com/pedroxer/resource-service/internal/config"
	"github.com/pedroxer/resource-service/internal/storage"
	log "github.com/sirupsen/logrus"
)

func main(){
	log := setupLogger()
	data, err := os.ReadFile("./configs/config.json")
	if err != nil {
		log.Fatal(err)
	}
	cfg := new(config.Config)

	if err := json.Unmarshal(data, cfg); err != nil {
		log.Fatal(err)
	}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}

	store, err := storage.NewStorage(&cfg.Postgres, log)
	if err != nil {
		log.Fatalf("failed connect to db %s", err)
	}
	log.Info("connected to db")
	
	app := app.NewApp(log, cfg.Port, store)
	if err := app.GRPCSrv.Run(); err != nil {
		log.Fatal(err)
	}
}

func setupLogger()*log.Logger{
	log := log.New()
	log.ReportCaller = true
	return log
}
