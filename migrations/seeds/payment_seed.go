package seeds

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"gorm.io/gorm"
)

type PaymentSeeder struct{}

func NewPaymentSeeder() *PaymentSeeder {
	return &PaymentSeeder{}
}

func (s *PaymentSeeder) Seed(db *gorm.DB) error {
	log.Println("[Seeder] Registrasi Payment dimulai... ")

	file, err := os.Open("./migrations/json/payments.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var paymentData []entity.Payment
	if err := json.NewDecoder(file).Decode(&paymentData); err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entity.Payment{}) {
		if err := db.Migrator().CreateTable(&entity.Payment{}); err != nil {
			return err
		}
	}

	for _, entry := range paymentData {
		var existing entity.Payment
		if err := db.First(&existing, "id = ?", entry.ID).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&entry).Error; err != nil {
				log.Printf("❌ Gagal insert Payment entry untuk user %s: %v\n", entry.UserID, err)
			} else {
				log.Printf("✅ Payment untuk %s berhasil: %s (%d rupiah)\n", entry.Type, entry.PaymentStatus, entry.PaidAmount)
			}
		} else {
			log.Printf("⚠️  Payment %s udah ada, ngapain diulang? 😑\n", entry.ID)
		}
	}

	log.Println("[Seeder] Registrasi Payment selesai. Jangan lupa bayar tagihan sendiri juga ya 🫠")
	return nil
}
