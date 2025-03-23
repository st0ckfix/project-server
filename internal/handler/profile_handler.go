package handler

import (
	"fastbuy/internal/auth"
	"fastbuy/internal/domain"
	"fastbuy/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileHandler struct {
	Usecase *usecase.ProfileUsecase
}

func NewProfileHandler(usecase *usecase.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{Usecase: usecase}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
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
	profile, err := h.Usecase.GetProfileBy(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": profile})
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
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
	var profile domain.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Usecase.UpdateProfile(username, &profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
