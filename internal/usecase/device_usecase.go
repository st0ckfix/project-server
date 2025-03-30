package usecase

import (
	"fastbuy/internal/domain"
	"fastbuy/internal/repository"
)

type DeviceUsecase struct {
	DeviceRepo *repository.DeviceRepository
}

func NewDeviceUsecase(repo *repository.DeviceRepository) *DeviceUsecase {
	return &DeviceUsecase{DeviceRepo: repo}
}

// Thêm thiết bị mới
func (uc *DeviceUsecase) AddDevice(username string, deviceName string, lat, lng float64) error {
	device := &domain.Device{
		DeviceName:  deviceName,
		Username:    username,
		Temperature: 0,
		Humidity:    0,
		Moisture:    0,
		Light:       0,
		Lat:         lat,
		Lng:         lng,
	}
	return uc.DeviceRepo.AddDevice(device)
}

// Lấy danh sách thiết bị
func (uc *DeviceUsecase) GetDevices(username string) ([]domain.Device, error) {
	return uc.DeviceRepo.GetDevices(username)
}

// Xóa thiết bị
func (uc *DeviceUsecase) RemoveDevice(deviceID uint) error {
	return uc.DeviceRepo.RemoveDevice(deviceID)
}

// Cập nhật thiết bị
func (uc *DeviceUsecase) UpdateDevice(deviceID uint, deviceName string, lat, lng float64) error {
	device := &domain.Device{
		DeviceID:   deviceID,
		DeviceName: deviceName,
		Lat:        lat,
		Lng:        lng,
	}
	return uc.DeviceRepo.UpdateDevice(device)
}
