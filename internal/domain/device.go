package domain

type Device struct {
	DeviceID    uint    `gorm:"primaryKey;autoIncrement"`
	Username    string  `gorm:"not null"` // Liên kết với user
	DeviceName  string  `gorm:"not null"`
	Temperature float64 `gorm:"default:0"`
	Humidity    int     `gorm:"default:0"`
	Moisture    int     `gorm:"default:0"`
	Light       int     `gorm:"default:0"`
	Lat         float64 `gorm:"not null"`
	Lng         float64 `gorm:"not null"`
}
