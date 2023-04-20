package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"notification_service/models"
	"notification_service/services"
)

type GetNotificationsRequest struct {
	UserId uint64
	Offset uint32
	Limit  uint32
}

type GetNotificationsResponse struct {
	Success       bool
	Notifications []models.Notification
	Total         uint32
}

func makeGetNotificationsEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNotificationsRequest)
		err, notifications, total := s.GetNotifications(
			req.UserId,
			req.Limit,
			req.Offset,
		)
		if err != nil {
			return GetNotificationsResponse{Success: false}, err
		}
		return GetNotificationsResponse{
			Success:       true,
			Notifications: notifications,
			Total:         total,
		}, nil
	}
}
