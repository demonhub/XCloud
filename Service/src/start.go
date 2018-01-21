package main

import (
	"github.com/gin-gonic/gin"
	"config"
)

func main() {
	r := gin.Default()

	initialBind(r)

	r.Run(config.ListenAddress) // listen and serve on 0.0.0.0:8080
}