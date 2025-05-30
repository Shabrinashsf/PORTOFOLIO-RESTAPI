package migrations

import (
	"log"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func Migrate() error {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err := initializers.DB.AutoMigrate(
		&entity.User{},
		&entity.SHSF{},
		&entity.Event{},
		&entity.Subevent{},
		&entity.EventPayment{},
		&entity.Payment{},
	); err != nil {
		log.Printf("Failed to migrate: %v", err)
		return err
	}

	log.Println("Database migration completed successfully.")
	return nil
}
