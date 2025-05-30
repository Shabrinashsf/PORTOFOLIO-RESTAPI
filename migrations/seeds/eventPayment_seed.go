package seeds

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"gorm.io/gorm"
)

type EventPaymentSeeder struct{}

func NewEventPaymentSeeder() *EventPaymentSeeder {
	return &EventPaymentSeeder{}
}

func (s *EventPaymentSeeder) Seed(db *gorm.DB) error {
	log.Println("[Seeder] Registrasi EventPayment dimulai... ")

	file, err := os.Open("./migrations/json/event_payments.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var data []entity.EventPayment
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entity.EventPayment{}) {
		if err := db.Migrator().CreateTable(&entity.EventPayment{}); err != nil {
			return err
		}
	}

	for _, entry := range data {
		var existing entity.EventPayment
		if err := db.First(&existing, "id = ?", entry.ID).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&entry).Error; err != nil {
				log.Printf("❌ Gagal insert EventPayment: Regist %s - Payment %s: %v\n", entry.RegistID, entry.PaymentID, err)
			} else {
				log.Printf("✅ EventPayment sukses: Regist %s → Payment %s\n", entry.RegistID, entry.PaymentID)
			}
		} else {
			log.Printf("⚠️  EventPayment %s udah ada... jangan maksa relasi yang udah pernah dicoba 🫠\n", entry.ID)
		}
	}

	log.Println("[Seeder] Registrasi EventPayment selesai. Fix jodohnya udah ketemu 😌")
	return nil
}
