package repository

import (
	"server/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGRepository struct {
	DB *gorm.DB
}

func NewPGRepository() *PGRepository {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5433 sslmode=disable"
	return getPGRepository(dsn)
}

func NewTestPGRepository() *PGRepository {
	dsn := "host=localhost user=postgres password=password dbname=test port=5433 sslmode=disable"
	return getPGRepository(dsn)
}

func getPGRepository(dsn string) *PGRepository {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
	db.AutoMigrate(&model.Product{}, &model.User{})
	return &PGRepository{DB: db}
}
