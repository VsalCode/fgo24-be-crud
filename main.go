package main

import (
	"crud-backend/routers"
	"crud-backend/utils"
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := utils.DBConnect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	r := gin.Default()
	routers.CombineRouter(r)

	log.Println("Server starting on port 8080...")
	r.Run(":8080")
}