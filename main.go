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

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			if err := migrations.Migrate(); err != nil {
				log.Fatalf("❌ Migration failed: %v", err)
			} else {
				log.Println("✅ Migration executed successfully.")
			}
			return

		case "seed":
			if err := migrations.RunSeeder(initializers.DB); err != nil {
				log.Fatalf("❌ Seeder failed: %v", err)
			} else {
				log.Println("✅ Seeder executed successfully.")
			}
			return
		}
	}

	var (
		// Implementation Dependency Injection
		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(initializers.DB)
		shsfRepository repository.SHSFRepository = repository.NewSHSFService(initializers.DB)

		// Service
		userService service.UserService = service.NewUserService(userRepository)
		shsfService service.SHSFService = service.NewSHSFService(shsfRepository)

		// Controller
		userController controllers.UserController = controllers.NewUserController(userService)
		shsfController controllers.SHSFController = controllers.NewSHSFController(shsfService)
	)

	r := gin.Default()
	routes.User(r, userController)
	routes.SHSF(r, shsfController)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
