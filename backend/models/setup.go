package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("campus_connect.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Mahasiswa{})
	DB.AutoMigrate(&Organisasi{})
	DB.AutoMigrate(&Universitas{})
	DB.AutoMigrate(&Fakultas{})
	DB.AutoMigrate(&Prodi{})
	DB.AutoMigrate(&Jabatan{})
	DB.AutoMigrate(&Post{})
	DB.AutoMigrate(&Comment{})
	DB.AutoMigrate(&Save{})
	DB.AutoMigrate(&Like{})
}
