package main

import (
	S "github.com/Zineb-ada/cyrkl/tree/social-framework/back/go-back-cyrkl/server"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
func main() {
	S.Server()
}
