package repository

import (
	_ "fastbuy/config"
	"fastbuy/internal/domain"
	_ "go.mongodb.org/mongo-driver/bson"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	DB *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{DB: db}
}

func (repo *ProfileRepository) GetProfile(username string) (*domain.Profile, error) {
	var profile domain.Profile
	err := repo.DB.Where("username = ?", username).First(&profile).Error
	return &profile, err
}

func (repo *ProfileRepository) UpdateProfile(username string, profile *domain.Profile) error {
	return repo.DB.Model(&domain.Profile{}).Where("username = ?", username).Updates(profile).Error
}
