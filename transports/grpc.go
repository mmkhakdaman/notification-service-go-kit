package transport

import (
	"notification_service/endpoints"
	"notification_service/pb"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
)

type gRPCServer struct {
	sendNotification     gt.Handler
	seenNotification     gt.Handler
	getNotifications     gt.Handler
	seenAllNotifications gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.NotificationServiceServer {
	return &gRPCServer{
		sendNotification: gt.NewServer(
			endpoints.SendNotification,
			decodeSendNotificationRequest,
			encodeSendNotificationResponse,
		),
		seenNotification: gt.NewServer(
			endpoints.SeenNotification,
			decodeSeenNotificationRequest,
			encodeSeenNotificationResponse,
		),
		getNotifications: gt.NewServer(
			endpoints.GetNotifications,
			decodeGetNotificationsRequest,
			encodeGetNotificationsResponse,
		),
		seenAllNotifications: gt.NewServer(
			endpoints.SeenAllNotifications,
			decodeSeenAllNotificationsRequest,
			encodeSeenAllNotificationsResponse,
		),
	}
}
