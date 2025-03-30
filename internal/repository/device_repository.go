package repository

import (
	"fastbuy/internal/domain"
	"gorm.io/gorm"
)

type DeviceRepository struct {
	DB *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{DB: db}
}

// Thêm thiết bị mới
func (repo *DeviceRepository) AddDevice(device *domain.Device) error {
	return repo.DB.Create(device).Error
}

// Lấy danh sách tất cả thiết bị
func (repo *DeviceRepository) GetDevices(username string) ([]domain.Device, error) {
	var devices []domain.Device
	err := repo.DB.Where("username = ?", username).Find(&devices).Error
	return devices, err
}

// Xóa thiết bị theo ID
func (repo *DeviceRepository) RemoveDevice(deviceID uint) error {
	return repo.DB.Delete(&domain.Device{}, "device_id = ?", deviceID).Error
}

// Cập nhật thông tin thiết bị
func (repo *DeviceRepository) UpdateDevice(device *domain.Device) error {
	return repo.DB.Model(&domain.Device{}).
		Where("device_id = ?", device.DeviceID).
		Updates(device).Error
}
