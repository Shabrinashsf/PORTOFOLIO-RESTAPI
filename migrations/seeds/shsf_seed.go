package seeds

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/entity"
	"gorm.io/gorm"
)

type SHSFSeeder struct{}

func NewSHSFSeeder() *SHSFSeeder {
	return &SHSFSeeder{}
}

func (s *SHSFSeeder) Seed(db *gorm.DB) error {
	log.Println("[Seeder] Registrasi SHSF dimulai... ")

	file, err := os.Open("migrations/json/shsfs.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var shsfData []entity.SHSF
	if err := json.NewDecoder(file).Decode(&shsfData); err != nil {
		return err
	}

	if !db.Migrator().HasTable(&entity.SHSF{}) {
		if err := db.Migrator().CreateTable(&entity.SHSF{}); err != nil {
			return err
		}
	}

	for _, entry := range shsfData {
		var existing entity.SHSF
		if err := db.First(&existing, "id = ?", entry.ID).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&entry).Error; err != nil {
				log.Printf("❌ Gagal insert SHSF entry %s (%s): %v\n", entry.Name, entry.Email, err)
			} else {
				log.Printf("✅ SHSF %s (%s) berhasil dimasukkan dengan status: %s\n", entry.Name, entry.Email, entry.AdminStatus)
			}
		} else {
			log.Printf("⚠️  SHSF %s udah ada, ga usah maksa ya...\n", entry.Email)
		}
	}

	log.Println("[Seeder] Registrasi SHSF selesai. Semoga ACC semua 🫠")
	return nil
}
