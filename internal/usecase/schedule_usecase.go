package usecase

import (
	"fastbuy/internal/domain"
	"fastbuy/internal/repository"
)

type ScheduleUsecase struct {
	ScheduleRepo *repository.ScheduleRepository
}

func NewScheduleUsecase(repo *repository.ScheduleRepository) *ScheduleUsecase {
	return &ScheduleUsecase{ScheduleRepo: repo}
}

// Lấy danh sách schedule của thiết bị
func (uc *ScheduleUsecase) GetSchedules(deviceID uint) ([]domain.Schedule, error) {
	return uc.ScheduleRepo.GetSchedules(deviceID)
}

// Thêm schedule
func (uc *ScheduleUsecase) AddSchedule(deviceId uint, description string, hour int, minute int, isRepeat bool, isSnooze bool, repeatList int) error {

	schedule := &domain.Schedule{
		DeviceID:    deviceId,
		Description: description,
		Hour:        hour,
		Minute:      minute,
		IsRepeat:    isRepeat,
		IsSnooze:    isSnooze,
		RepeatList:  repeatList,
	}
	return uc.ScheduleRepo.AddSchedule(schedule)
}

// Cập nhật schedule
func (uc *ScheduleUsecase) UpdateSchedule(schedule *domain.Schedule) error {
	return uc.ScheduleRepo.UpdateSchedule(schedule)
}

// Xóa schedule
func (uc *ScheduleUsecase) RemoveSchedule(scheduleID uint) error {
	return uc.ScheduleRepo.RemoveSchedule(scheduleID)
}
