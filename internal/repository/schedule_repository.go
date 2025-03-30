package repository

import (
	"fastbuy/internal/domain"
	"gorm.io/gorm"
)

type ScheduleRepository struct {
	DB *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) *ScheduleRepository {
	return &ScheduleRepository{DB: db}
}

// Lấy tất cả schedules của một thiết bị
func (repo *ScheduleRepository) GetSchedules(deviceID uint) ([]domain.Schedule, error) {
	var schedules []domain.Schedule
	err := repo.DB.Where("device_id = ?", deviceID).Find(&schedules).Error
	return schedules, err
}

// Thêm mới schedule
func (repo *ScheduleRepository) AddSchedule(schedule *domain.Schedule) error {
	return repo.DB.Create(schedule).Error
}

// Cập nhật schedule
func (repo *ScheduleRepository) UpdateSchedule(schedule *domain.Schedule) error {
	return repo.DB.Save(schedule).Error
}

// Xóa schedule
func (repo *ScheduleRepository) RemoveSchedule(scheduleID uint) error {
	return repo.DB.Delete(&domain.Schedule{}, scheduleID).Error
}
