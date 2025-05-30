package seeds

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"gorm.io/gorm"
)

type EventSeeder struct{}

func NewEventSeeder() *EventSeeder {
	return &EventSeeder{}
}

func (s *EventSeeder) Seed(db *gorm.DB) error {
	log.Println("[Seeder] Registrasi Event dimulai... ")

	file, err := os.Open("./migrations/json/events.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var eventData []entity.Event
	if err := json.NewDecoder(file).Decode(&eventData); err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entity.Event{}) {
		if err := db.Migrator().CreateTable(&entity.Event{}); err != nil {
			return err
		}
	}

	for _, entry := range eventData {
		var existing entity.Event
		if err := db.First(&existing, "id = ?", entry.ID).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&entry).Error; err != nil {
				log.Printf("❌ Gagal insert Event %s: %v\n", entry.Name, err)
			} else {
				log.Printf("✅ Event %s berhasil dimasukkan ke timeline sejarah 🎉\n", entry.Name)
			}
		} else {
			log.Printf("⚠️  Event %s udah ada... masa mau diulang terus kayak mantan? 😬\n", entry.Name)
		}
	}

	log.Println("[Seeder] Registrasi Event selesai. Jalanin event-nya jangan pasang wajah lesu ya 🙃")
	return nil
}
