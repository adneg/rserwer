package main

import (
	"dbtools"
	//"log"
	"restapi"
)

func main() {
	dbtools.Start()
	restapi.Start()

}
