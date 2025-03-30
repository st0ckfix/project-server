package handler

import (
	"fastbuy/internal/domain"
	"fastbuy/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ScheduleHandler struct {
	Usecase *usecase.ScheduleUsecase
}

func NewScheduleHandler(uc *usecase.ScheduleUsecase) *ScheduleHandler {
	return &ScheduleHandler{Usecase: uc}
}

// Lấy danh sách schedule của một thiết bị
func (h *ScheduleHandler) GetSchedules(c *gin.Context) {
	deviceID, err := strconv.Atoi(c.Query("device_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device ID"})
		return
	}

	schedules, err := h.Usecase.GetSchedules(uint(deviceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get schedules"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}

// Thêm schedule mới
func (h *ScheduleHandler) AddSchedule(c *gin.Context) {

	var request struct {
		DeviceId    uint   `json:"device_id"`
		Description string `json:"description"`
		Hour        int    `json:"hour"`
		Minute      int    `json:"minute"`
		IsRepeat    bool   `json:"is_repeat"`
		IsSnooze    bool   `json:"is_snooze"`
		RepeatList  int    `json:"repeat_list"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.Usecase.AddSchedule(request.DeviceId, request.Description, request.Hour, request.Minute, request.IsRepeat, request.IsSnooze, request.RepeatList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule added successfully"})
}

// Cập nhật schedule
func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	var schedule domain.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.Usecase.UpdateSchedule(&schedule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule updated successfully"})
}

// Xóa schedule
func (h *ScheduleHandler) RemoveSchedule(c *gin.Context) {
	scheduleID, err := strconv.Atoi(c.Query("schedule_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	if err := h.Usecase.RemoveSchedule(uint(scheduleID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule removed successfully"})
}
