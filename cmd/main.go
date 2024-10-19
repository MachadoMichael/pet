package main

import (
	"github.com/MachadoMichael/pet/api/server"
	"github.com/MachadoMichael/pet/config"
)

func main() {
	config.Load()
	server.Start()
}
