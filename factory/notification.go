package factory

import (
	"github.com/brianvoe/gofakeit/v6"
	"notification_service/models"
)

// Factory is the interface for the factory.
type Factory interface {
	// CreateNotification creates a new notification.
	CreateNotification() *models.Notification
}

// FactoryImpl is the implementation of the Factory interface.
type FactoryImpl struct {
}

// NewFactory returns a new instance of FactoryImpl.
func NewFactory() *FactoryImpl {
	return &FactoryImpl{}
}

// CreateNotification creates a new notification.
func (f *FactoryImpl) CreateNotification() *models.Notification {
	return &models.Notification{
		Title:            gofakeit.Name(),
		Message:          gofakeit.Name(),
		UserId:           gofakeit.Uint64(),
		SenderId:         gofakeit.Uint64(),
		NotificationType: gofakeit.Name(),
		Avatar:           gofakeit.Name(),
		Link:             gofakeit.Name(),
		LinkText:         gofakeit.Name(),
		Icon:             gofakeit.Name(),
	}
}
