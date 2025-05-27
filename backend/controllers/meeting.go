package controllers

import (
	"FullStack/config"
	"FullStack/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateMeeting(c *gin.Context) {
	var meeting models.Meeting
	if err := c.ShouldBindJSON(&meeting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)
	meeting.OwnerID = userID
	config.DB.Create(&meeting)

	c.JSON(http.StatusOK, meeting)
}

func GetMeetings(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var meetings []models.Meeting
	config.DB.Where("owner_id = ?", userID).Find(&meetings)

	c.JSON(http.StatusOK, meetings)
}

func UpdateMeeting(c *gin.Context) {
	id := c.Param("id")
	var meeting models.Meeting
	if err := config.DB.First(&meeting, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "جلسه پیدا نشد"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if meeting.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "اجازه دسترسی ندارید"})
		return
	}

	var input models.Meeting
	c.ShouldBindJSON(&input)

	meeting.Title = input.Title
	meeting.DateTime = input.DateTime

	config.DB.Save(&meeting)
	c.JSON(http.StatusOK, meeting)
}

func DeleteMeeting(c *gin.Context) {
	id := c.Param("id")
	var meeting models.Meeting
	if err := config.DB.First(&meeting, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "جلسه پیدا نشد"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if meeting.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "اجازه ندارید"})
		return
	}

	config.DB.Delete(&meeting)
	c.JSON(http.StatusOK, gin.H{"message": "جلسه حذف شد"})
}
func InviteUser(c *gin.Context) {
	meetingID := c.Param("id")
	var meeting models.Meeting
	if err := config.DB.First(&meeting, meetingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "جلسه پیدا نشد"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if meeting.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "فقط صاحب جلسه می‌تونه دعوت کنه"})
		return
	}

	var input struct {
		UserID uint `json:"user_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invitation := models.MeetingUser{
		UserID:    input.UserID,
		MeetingID: meeting.ID,
	}
	config.DB.Create(&invitation)
	c.JSON(http.StatusOK, gin.H{"message": "کاربر دعوت شد"})
}
