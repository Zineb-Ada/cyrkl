package server

import (
	"fmt"
	"os"

	"github.com/Zineb-ada/cyrkl/tree/social-framework/back/go-back-cyrkl/user"
	"github.com/gin-gonic/gin"
)

/*
 Runs our server and houses all of our route endpoints
*/
func Server() {
	if os.Getenv("BUILD") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("BUILD: " + os.Getenv("BUILD"))
		fmt.Println("PORT: " + os.Getenv("PORT"))
	}
	router := gin.Default()
	user.Routes(router)
	router.Run(":" + os.Getenv("PORT"))
}
