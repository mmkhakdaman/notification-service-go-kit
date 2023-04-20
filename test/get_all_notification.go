package test

import (
	"github.com/brianvoe/gofakeit/v6"
	"notification_service/database"
	"notification_service/models"
	"notification_service/services"
	"testing"
)

// TestGetAllNotifications tests the GetAllNotifications method.
func TestGetAllNotifications(t *testing.T) {
	// Create a new service.
	s := services.NewService(log.NewNopLogger(), database.GetDB())

	// Create a new user.
	user := &models.User{
		Username: gofakeit.Username(),
	}

	// Save the user to the database.
	err := s.SaveUser(user)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Create a new notification.
	notification := &models.Notification{
		Title:            gofakeit.Name(),
		Message:          gofakeit.Name(),
		UserId:           user.ID,
		SenderId:         user.ID,
		NotificationType: gofakeit.Name(),
		Avatar:           gofakeit.Name(),
		Link:             gofakeit.Name(),
		LinkText:         gofakeit.Name(),
		Icon:             gofakeit.Name(),
	}

}
