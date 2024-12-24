package main

import (
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/migrations"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := migrations.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		} else {
			log.Println("Migration executed successfully.")
		}
	} else {
		log.Println("No valid argument provided. Usage: `go run main.go migrate`")
	}

	r.Run()
}
