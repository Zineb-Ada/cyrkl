package main

import (
	"github.com/joho/godotenv"
	S "github.com/zineb-Ada/cyrkl/back/go-back-cyrkl/server"
)

func init() {
	godotenv.Load()
}
func main() {
	S.Server()
}
