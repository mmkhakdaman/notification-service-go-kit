package test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"notification_service/database"
	"notification_service/models"
	"notification_service/services"
	"testing"

	"github.com/go-kit/kit/log"
)

func TestSendNotification(t *testing.T) {
	db, cleanup := database.CreateTestGormDB(t)
	defer cleanup()

	s := services.NewService(log.NewNopLogger(), database.GetDB())

	gofakeit.Seed(0)
	title := gofakeit.Name()
	message := gofakeit.Name()
	recipient := gofakeit.Uint64()
	sender := gofakeit.Uint64()
	notificationType := gofakeit.Name()
	avatar := gofakeit.Name()
	link := gofakeit.Name()
	linkText := gofakeit.Name()
	icon := gofakeit.Name()

	// Create a notification.
	err := s.SendNotification(
		title,
		message,
		recipient,
		sender,
		notificationType,
		avatar,
		link,
		linkText,
		icon,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check that the notification was created.
	var notification models.Notification
	err = db.First(&notification).Error

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// assert that the notification was created correctly.
	assert.Equal(t, title, notification.Title)
	assert.Equal(t, message, notification.Message)
	assert.Equal(t, recipient, notification.UserId)
	assert.Equal(t, sender, notification.SenderId)
	assert.Equal(t, notificationType, notification.NotificationType)
	assert.Equal(t, avatar, notification.Avatar)
	assert.Equal(t, link, notification.Link)
	assert.Equal(t, linkText, notification.LinkText)
	assert.Equal(t, icon, notification.Icon)
}
