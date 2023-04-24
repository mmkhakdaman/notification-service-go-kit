package test

import (
	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"notification_service/database"
	"notification_service/factory"
	"notification_service/models"
	"notification_service/services"
	"sort"
	"testing"
)

// TestGetAllNotifications tests the GetAllNotifications method.
func TestGetAllNotifications(t *testing.T) {
	gormDB, cleanup := database.CreateTestGormDB(t)
	defer cleanup()

	// notification factory
	factoryService := factory.NewFactory()

	var notifications []*models.Notification
	// create a notification
	for i := 0; i < 10; i++ {
		notif := factoryService.CreateNotification()
		notif.UserId = 50
		notif.ID = uint64(i + 1)
		notifications = append(notifications, notif)
	}

	gormDB.Create(&notifications)

	// get all notifications
	// assert that the notifications are the same
	service := services.NewService(log.NewNopLogger(), database.GetDB())
	fetchedNotification, total, err := service.GetNotifications(uint64(50), uint32(10), uint32(0))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	sort.Slice(notifications, func(i, j int) bool {
		return notifications[i].ID > notifications[j].ID
	})

	// assert that the notifications are the same
	assert.EqualValues(t, total, 10)
	for i := 10; i > 10; i-- {
		assert.EqualValues(t, notifications[i].ID, fetchedNotification[i].ID)
	}
}
