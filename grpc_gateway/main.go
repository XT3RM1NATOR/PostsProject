package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()

	//routes.SetupRoutes(r, db)

	if err := r.Run(); err != nil {
		log.Fatal("❌Failed starting the server: ❌", err)
	}

}
