package repository

import (
	"context"
	"time"

	_ "fastbuy/config"
	"fastbuy/internal/domain"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *domain.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) FindUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	err := repo.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

type MongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(client *mongo.Client) *MongoRepository {
	return &MongoRepository{Collection: client.Database("app_db").Collection("login_logs")}
}

func (repo *MongoRepository) LogLogin(username string) error {
	_, err := repo.Collection.InsertOne(context.TODO(), domain.UserLoginLog{
		Username:  username,
		Timestamp: time.Now(),
	})
	return err
}
