package usecase

import (
	"fastbuy/internal/domain"
	"fastbuy/internal/repository"
)

type ProfileUsecase struct {
	ProfileRepo *repository.ProfileRepository
}

func NewProfileUsecase(profileRepo *repository.ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{ProfileRepo: profileRepo}
}

func (uc *ProfileUsecase) GetProfileBy(username string) (*domain.Profile, error) {
	profile, err := uc.ProfileRepo.GetProfile(username)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (uc *ProfileUsecase) UpdateProfile(username string, profile *domain.Profile) error {
	err := uc.ProfileRepo.UpdateProfile(username, profile)
	if err != nil {
		return err
	}
	return nil
}
