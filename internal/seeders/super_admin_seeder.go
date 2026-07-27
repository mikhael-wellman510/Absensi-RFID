package seeders

import (
	"attendance-api/internal/entities"
	"attendance-api/internal/enums"
	"attendance-api/internal/utils"
	"log"

	"gorm.io/gorm"
)

func SeedSuperAdmin(db *gorm.DB) {
	var count int64
	// Mengecek apakah sudah ada user dengan role SuperAdmin di database
	err := db.Model(&entities.User{}).Where("role = ?", enums.SuperAdmin).Count(&count).Error
	if err != nil {
		log.Printf("Gagal mengecek keberadaan Super Admin: %v", err)
		return
	}

	// Jika belum ada SuperAdmin sama sekali, buat 1 akun default
	if count == 0 {
		hashedPassword, err := utils.HashPassword("SuperSecret123!") // Ganti dengan password yang aman
		if err != nil {
			log.Printf("Gagal melakukan hash password seeder: %v", err)
			return
		}

		admin1 := &entities.User{
			FullName:        "System Administrator1",
			Email:           "superadmin1@attendance.com", // Email default untuk login
			PhoneNumber:     "000000000000",
			Password:        hashedPassword,
			Role:            enums.SuperAdmin,
			IsActive:        true,
			IsEmailVerified: true,
		}
		admin2 := &entities.User{
			FullName:        "System Administrator2",
			Email:           "superadmin2@attendance.com", // Email default untuk login
			PhoneNumber:     "000000000000",
			Password:        hashedPassword,
			Role:            enums.SuperAdmin,
			IsActive:        true,
			IsEmailVerified: true,
		}

		var admins []*entities.User

		listAdmin := append(admins, admin1, admin2)

		if err := db.Create(listAdmin).Error; err != nil {
			log.Fatalf("Gagal melakukan seeding Super Admin: %v", err)
		}

		log.Println("Seeding berhasil: Default Super Admin telah dibuat.")
	} else {
		log.Println("Seeding diabaikan: Super Admin sudah ada di database.")
	}
}
