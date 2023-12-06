package handler

import (
	"log"
	"tv360/src/db"
	"tv360/src/util"

	"github.com/robfig/cron/v3"
)

type Server struct {
	cfg     util.Config
	store   db.Store
	cronjob *cron.Cron
}

func NewServer(config util.Config, store db.Store) (*Server, error) {

	cronjob := cron.New()
	server := &Server{cfg: config, store: store, cronjob: cronjob}

	return server, nil
}

func (server *Server) Start() error {

	err := server.registerJob()
	if err != nil {
		log.Printf("can not register job: %s", err)
	}
	server.cronjob.Start()
	var forever chan struct{}
	<-forever
	return nil
}
