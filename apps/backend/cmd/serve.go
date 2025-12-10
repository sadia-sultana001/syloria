package cmd

import (
	"syloria-demo/config"
	"syloria-demo/rest"
)

func Serve() {

	cnf := config.GetConfig()

	rest.Start(cnf)
}
