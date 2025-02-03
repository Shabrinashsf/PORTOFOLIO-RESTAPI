package main

import (
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/controllers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/migrations"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/repository"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/routes"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/service"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	utils.LoadTemplates("template")

	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := migrations.Migrate(); err != nil {
			log.Fatalf("Migration failed: %v", err)
		} else {
			log.Println("Migration executed successfully.")
		}
		return
	}

	var (
		// Implementation Dependency Injection
		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(initializers.DB)

		// Service
		userService service.UserService = service.NewUserService(userRepository)

		// Controller
		userController controllers.UserController = controllers.NewUserController(userService)
	)

	r := gin.Default()
	routes.User(r, userController)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
