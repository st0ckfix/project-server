package domain

// Profile model for PostgreSQL
type Profile struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Firstname string
	Lastname  string
	Avatar    string
	Birthdate string
}
