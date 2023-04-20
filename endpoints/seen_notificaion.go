package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"notification_service/services"
)

// SeenNotificationRequest MathReq struct holds the request parameters for the Add method
type SeenNotificationRequest struct {
	NotificationId uint64
}

// SeenNotificationResponse MathResp struct holds the response parameters for the Add method
type SeenNotificationResponse struct {
	Success bool
}

// makeSeenNotificationEndpoint func initializes the Add endpoint
func makeSeenNotificationEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SeenNotificationRequest)
		err := s.SeenNotification(
			req.NotificationId,
		)
		if err != nil {
			return SeenNotificationResponse{Success: false}, err
		}
		return SeenNotificationResponse{Success: true}, nil
	}
}
