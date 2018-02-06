package main

import (
	"github.com/gin-gonic/gin"
	"config"
)

func main() {
	r := gin.Default()

	// Bind router
	initialBind(r)

	// listen and serve
	r.Run(config.ListenAddress)
}