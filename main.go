package main

import (
	"github.com/lucas-10101/logapi/server"
	"github.com/lucas-10101/logapi/settings"
)

func main() {
	settings.Configure()
	server.InitServer()
}
