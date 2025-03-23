package usecase

import (
	"errors"
	"fastbuy/internal/auth"
	"fastbuy/internal/domain"
	"fastbuy/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepo *repository.UserRepository
	LogRepo  *repository.MongoRepository
}

func NewUserUsecase(userRepo *repository.UserRepository, logRepo *repository.MongoRepository) *UserUsecase {
	return &UserUsecase{UserRepo: userRepo, LogRepo: logRepo}
}

func (uc *UserUsecase) Register(user *domain.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	return uc.UserRepo.CreateUser(user)
}

func (uc *UserUsecase) Login(username, password string) (*domain.User, string, string, error) {
	user, err := uc.UserRepo.FindUserByUsername(username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, "", "", nil
	}

	// Tạo accessToken & refreshToken
	accessToken := ""
	accessToken, _ = auth.GenerateAccessToken(user.Username)

	refreshToken := ""
	refreshToken, _ = auth.GenerateRefreshToken(user.Username)

	// Log login in MongoDB
	_ = uc.LogRepo.LogLogin(username)
	return user, accessToken, refreshToken, nil
}

// Logout blacklists both access and refresh tokens
func (uc *UserUsecase) Logout(accessToken, refreshToken string) error {
	// Blacklist both tokens
	err1 := auth.BlacklistToken(accessToken)
	err2 := auth.BlacklistToken(refreshToken)

	// Return error if either operation fails
	if err1 != nil {
		return err1
	}
	return err2
}

func (uc *UserUsecase) RefreshToken(refreshToken string) (string, error) {
	claims, status := auth.ParseToken(refreshToken)
	if status != auth.TOKEN_VALID {
		return "", errors.New("invalid refresh token")
	}

	username := claims["username"].(string)

	// Blacklist the used refresh token
	_ = auth.BlacklistToken(refreshToken)

	// Generate new access token
	newAccessToken, err := auth.GenerateAccessToken(username)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
