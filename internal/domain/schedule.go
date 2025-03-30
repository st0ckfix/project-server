package domain

type Schedule struct {
	ScheduleID  uint   `gorm:"primaryKey;autoIncrement"`
	DeviceID    uint   `gorm:"not null"`
	Description string `gorm:"type:text"`
	Hour        int    `gorm:"check:hour >= 0 AND hour < 24"`
	Minute      int    `gorm:"check:minute >= 0 AND minute < 60"`
	IsRepeat    bool   `gorm:"default:false"`
	IsSnooze    bool   `gorm:"default:false"`
	RepeatList  int    `gorm:"type:json"`
}
