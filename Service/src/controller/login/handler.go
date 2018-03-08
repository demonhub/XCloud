package login

import (
	"github.com/gin-gonic/gin"
	"model"
	"db"
	"fmt"
	"crypto/md5"
	"crypto"
	"crypto/sha1"
	"crypto/sha256"
	"util"
)

func Handler(c *gin.Context) {
	username := c.PostForm("U")
	password := c.PostForm("P")
	// try login
	res,ok := getLogin(username,password)
	if ok {
		c.JSON(200,gin.H{

		})
	} else {
		c.JSON(200,util.GetErrorJSON(model.LoginErr))
	}
}

func getLogin(u, p string) (model.User,bool) {
	// sha256 twice
	SHA256 := sha256.New()
	temp := SHA256.Sum([]byte(p))
	p = string(SHA256.Sum(temp))

	// SELECT database
	var DB = db.DB.Conn
	target := model.User{}
	err := DB.Get(&target, "SELECT * FROM USER WHERE USERNAME=$1 AND PASSWORD=$2", u,p)
	if err != nil {
		return target,false
	}
	return target,true
}