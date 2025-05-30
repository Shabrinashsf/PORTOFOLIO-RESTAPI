package seeds

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"gorm.io/gorm"
)

type SubeventSeeder struct{}

func NewSubeventSeeder() *SubeventSeeder {
	return &SubeventSeeder{}
}

func (s *SubeventSeeder) Seed(db *gorm.DB) error {
	log.Println("[Seeder] Seeding subevent dimulai...")

	file, err := os.Open("./migrations/json/subevents.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var subevents []entity.Subevent
	if err := json.NewDecoder(file).Decode(&subevents); err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entity.Subevent{}) {
		if err := db.Migrator().CreateTable(&entity.Subevent{}); err != nil {
			return err
		}
	}

	for _, sub := range subevents {
		var existing entity.Subevent
		if err := db.First(&existing, "id = ?", sub.ID).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&sub).Error; err != nil {
				log.Printf("❌ Gagal insert subevent %s: %v\n", sub.Name, err)
			} else {
				log.Printf("✅ Subevent %s berhasil ditambahkan!\n", sub.Name)
			}
		} else {
			log.Printf("⚠️  Subevent %s sudah ada, skip...\n", sub.Name)
		}
	}

	log.Println("[Seeder] Subevent selesai dengan selamat.")
	return nil
}
