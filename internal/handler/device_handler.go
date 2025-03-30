package handler

import (
	"fastbuy/internal/auth"
	"fastbuy/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeviceHandler struct {
	Usecase *usecase.DeviceUsecase
}

func NewDeviceHandler(uc *usecase.DeviceUsecase) *DeviceHandler {
	return &DeviceHandler{Usecase: uc}
}

// API Thêm thiết bị mới
func (h *DeviceHandler) AddDevice(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
		return
	}

	// Remove "Bearer " prefix
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, status := auth.ParseToken(tokenString)
	if status != auth.TOKEN_VALID {
		c.JSON(http.StatusUnauthorized, gin.H{"status": status, "error": "Invalid token"})
		return
	}

	username := claims["username"].(string)

	var request struct {
		DeviceName string  `json:"device_name"`
		Lat        float64 `json:"lat"`
		Lng        float64 `json:"lng"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Usecase.AddDevice(username, request.DeviceName, request.Lat, request.Lng)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add device"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Device added successfully"})
}

// API Lấy danh sách thiết bị
func (h *DeviceHandler) GetDevices(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
		return
	}

	// Remove "Bearer " prefix
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	claims, status := auth.ParseToken(tokenString)
	if status != auth.TOKEN_VALID {
		c.JSON(http.StatusUnauthorized, gin.H{"status": status, "error": "Invalid token"})
		return
	}

	username := claims["username"].(string)

	devices, err := h.Usecase.GetDevices(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get devices"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"devices": devices})
}

// API Xóa thiết bị
func (h *DeviceHandler) RemoveDevice(c *gin.Context) {
	var request struct {
		DeviceID uint `json:"device_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Usecase.RemoveDevice(request.DeviceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove device"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Device removed successfully"})
}

// API Cập nhật thiết bị
func (h *DeviceHandler) UpdateDevice(c *gin.Context) {
	var request struct {
		DeviceID   uint    `json:"device_id"`
		DeviceName string  `json:"device_name"`
		Lat        float64 `json:"lat"`
		Lng        float64 `json:"lng"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Usecase.UpdateDevice(request.DeviceID, request.DeviceName, request.Lat, request.Lng)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update device"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Device updated successfully"})
}
