package main

import (
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/migrations"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := migrations.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		} else {
			log.Println("Migration executed successfully.")
		}
		return
	}

	r := gin.Default()

	routes.User(r)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
