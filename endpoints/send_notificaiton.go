package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"notification_service/services"
)

// SendNotificationRequest MathReq struct holds the request parameters for the Add method
type SendNotificationRequest struct {
	Title     string
	Message   string
	Recipient uint64
	Sender    uint64
	Type      string
	Avatar    string
	Link      string
	LinkText  string
	Icon      string
}

// SendNotificationResponse MathResp struct holds the response parameters for the Add method
type SendNotificationResponse struct {
	Success bool
}

// makeSendNotificationEndpoint func initializes the Add endpoint
func makeSendNotificationEndpoint(s services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendNotificationRequest)
		err := s.SendNotification(
			req.Title,
			req.Message,
			req.Recipient,
			req.Sender,
			req.Type,
			req.Avatar,
			req.Link,
			req.LinkText,
			req.Icon,
		)
		if err != nil {
			return SendNotificationResponse{Success: false}, err
		}
		return SendNotificationResponse{Success: true}, nil
	}
}
