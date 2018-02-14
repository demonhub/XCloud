package main

import (
	"github.com/gin-gonic/gin"
	"config"
	"db"
)

func main() {
	// Connect database
	db.InitDatabase()

	r := gin.Default()

	// Bind router
	initialBind(r)

	// listen and serve
	r.Run(config.ListenAddress)
}