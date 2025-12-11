package main

import (
	"fmt"
	"syloria-demo/util"
)

//"syloria-demo/cmd"

func main() {
	//cnf := config.GetConfig()

	//fmt.Println(cnf.Version)
	//fmt.Println(cnf.ServiceName)
	//fmt.Println(cnf.HttpPort)
	//	cmd.Serve()

	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         43,
		FirstName:   "Sadia",
		LastName:    "Sultana",
		Email:       "sss@gmail.com",
		IsShopOwner: false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(jwt)

}
