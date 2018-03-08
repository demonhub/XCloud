package util

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
)

func GetOKJSON(parms map[string]string) gin.H {
	return gin.H{
		"msg": "",
		"code": 1,
	}
}

func GetErrorJSON(err error) gin.H {
	return gin.H{
		"msg": err.Error(),
		"code": -1,
	}
}