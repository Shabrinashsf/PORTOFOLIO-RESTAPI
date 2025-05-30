package migrations

import (
	"log"

	"github.com/Shabrinashsf/PORTOFOLIO-RESTAPI/migrations/seeds"
	"gorm.io/gorm"
)

type Seeder interface {
	Seed(db *gorm.DB) error
}

func RunSeeder(db *gorm.DB) error {
	log.Println("🚜 [Seeder] Proses seeding database dimulai...")

	seeders := []Seeder{
		seeds.NewUserSeeder(),
		seeds.NewEventSeeder(),
		seeds.NewSubeventSeeder(),
		seeds.NewSHSFSeeder(),
		seeds.NewPaymentSeeder(),
		seeds.NewEventPaymentSeeder(),
	}

	for _, seeder := range seeders {
		log.Printf("🔧 [Seeder] Menjalankan: %T", seeder)
		if err := seeder.Seed(db); err != nil {
			log.Printf("❌ [Seeder] Gagal di %T: %v", seeder, err)
			return err
		}
	}

	log.Println("✅ [Seeder] Semua seeder berhasil dijalankan.")
	return nil
}
