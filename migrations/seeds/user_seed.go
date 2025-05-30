package seeds

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func NewUserSeeder() *UserSeeder {
	return &UserSeeder{}
}

func (s *UserSeeder) Seed(db *gorm.DB) error {
	log.Println("[Seeder] Seeding user data dimulai...")

	file, err := os.Open("./migrations/json/users.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var users []entity.User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entity.User{}) {
		if err := db.Migrator().CreateTable(&entity.User{}); err != nil {
			return err
		}
	}

	for _, u := range users {
		var existing entity.User
		if err := db.First(&existing, "email = ?", u.Email).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&u).Error; err != nil {
				log.Printf("Gagal menambahkan user %s: %v\n", u.Email, err)
				continue
			}
			log.Printf("User %s berhasil ditambahkan!\n", u.Email)
		} else {
			log.Printf("User %s sudah ada, skip...\n", u.Email)
		}
	}

	log.Println("[Seeder] User seeding selesai.")
	return nil
}
