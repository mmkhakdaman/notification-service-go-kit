package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"notification_service/services"
)

type SeenAllNotificationsRequest struct {
	UserId uint64
}

type SeenAllNotificationsResponse struct {
	Success bool
}

func makeSeenAllNotificationsEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SeenAllNotificationsRequest)
		err := s.SeenAllNotifications(
			req.UserId,
		)
		if err != nil {
			return GetNotificationsResponse{Success: false}, err
		}
		return SeenAllNotificationsResponse{
			Success: true,
		}, nil
	}
}
