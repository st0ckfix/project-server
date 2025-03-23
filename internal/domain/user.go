package domain

import "time"

// User model for PostgreSQL
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
}

// UserLoginLog model for MongoDB
type UserLoginLog struct {
	ID        string    `bson:"_id,omitempty"`
	Username  string    `bson:"username"`
	Timestamp time.Time `bson:"timestamp"`
}
