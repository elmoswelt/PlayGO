package main

import (
	"fmt"
	 "github/elmoswelt/Geekon-2016/Source/server"
)

func main() {

	fmt.Println("*** Geekon 2016 Localization Service *** \n")

	myServer := new(server.Server)
	myServer.Start()
}


