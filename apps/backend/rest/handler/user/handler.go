package user

import (
	"syloria-demo/config"
)

type Handler struct {
	cnf *config.Config
	svr Service
}

func NewHandler(cnf *config.Config, svr Service) *Handler {
	return &Handler{
		cnf: cnf,
		svr: svr,
	}
}
