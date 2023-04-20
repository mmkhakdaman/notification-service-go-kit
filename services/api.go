package services

import (
	"notification_service/database"
	"notification_service/models"

	"gorm.io/gorm"

	"github.com/go-kit/kit/log"
)

type service struct {
	logger log.Logger
	db     *gorm.DB
}

// Service interface describes a service that adds numbers
type Service interface {
	SendNotification(title string,
		message string,
		recipient uint64,
		sender uint64,
		notificationType string,
		avatar string,
		link string,
		linkText string,
		icon string) error
	SeenNotification(id uint64) error
	SeenAllNotifications(id uint64) error
	GetNotifications(userId uint64, limit uint32, offset uint32) (error, []models.Notification, uint32)
}

// NewService returns a Service with all the expected dependencies
func NewService(logger log.Logger, db *gorm.DB) Service {
	return &service{
		logger: logger,
		db:     db,
	}
}

// SendNotification SaveNotification saves the notification to the repository
func (r *service) SendNotification(
	title string,
	message string,
	recipient uint64,
	sender uint64,
	notificationType string,
	avatar string,
	link string,
	linkText string,
	icon string,
) error {
	// implementation details for saving the notification, such as interacting with a database
	notification := &models.Notification{
		Title:            title,
		SenderId:         sender,
		UserId:           recipient,
		Message:          message,
		Avatar:           avatar,
		Link:             link,
		LinkText:         linkText,
		Icon:             icon,
		NotificationType: notificationType,
	}

	return r.db.Create(notification).Error
}

func (r *service) SeenNotification(id uint64) error {
	db := database.GetDB()

	notification := &models.Notification{
		ID: id,
	}

	return db.Model(notification).Update("seen", true).Error
}

func (r *service) GetNotifications(id uint64, limit uint32, offset uint32) (error, []models.Notification, uint32) {
	db := database.GetDB()

	var notifications []models.Notification
	var count int64

	err := db.Model(&models.Notification{}).Where("user_id = ?", id).Count(&count).Error
	if err != nil {
		return err, nil, 0
	}

	err = db.Model(&models.Notification{}).Where("user_id = ?", id).Order("created_at desc").Limit(int(limit)).Offset(int(offset)).Find(&notifications).Error
	if err != nil {
		return err, nil, 0
	}

	return nil, notifications, uint32(count)
}

func (r *service) SeenAllNotifications(id uint64) error {
	db := database.GetDB()

	notification := &models.Notification{
		UserId: id,
	}

	return db.Model(notification).Update("seen", true).Error
}
