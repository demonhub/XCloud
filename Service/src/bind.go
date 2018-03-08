package main

import "controller/login"
import "github.com/gin-gonic/gin"

func initialBind(r *gin.Engine) {
	r.GET("/login",login.Handler)
}
