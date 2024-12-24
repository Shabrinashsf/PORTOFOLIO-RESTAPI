package migrations

import (
	"log"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func Migrate() error {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := initializers.DB.AutoMigrate(
		&models.User{},
	); err != nil {
		log.Printf("Failed to migrate: %v", err)
		return err
	}

	log.Println("Database migration completed successfully.")
	return nil
}
