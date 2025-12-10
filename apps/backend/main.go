package main

import (
	"syloria-demo/cmd"
)

func main() {
	//cnf := config.GetConfig()

	//fmt.Println(cnf.Version)
	//fmt.Println(cnf.ServiceName)
	//fmt.Println(cnf.HttpPort)
	cmd.Serve()

}
