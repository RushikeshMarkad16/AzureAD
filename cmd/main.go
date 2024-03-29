package main

import (
	"github.com/RushikeshMarkad16/AzureAD/config"
	"github.com/RushikeshMarkad16/AzureAD/server"
)

func main() {
	config.Load()
	server.StartServer()
}
