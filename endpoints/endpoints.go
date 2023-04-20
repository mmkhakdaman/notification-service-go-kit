package endpoints

import (
	"notification_service/services"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints struct holds the list of endpoints definition
type Endpoints struct {
	SendNotification     endpoint.Endpoint
	SeenNotification     endpoint.Endpoint
	GetNotifications     endpoint.Endpoint
	SeenAllNotifications endpoint.Endpoint
}

// MakeEndpoints func initializes the Endpoint instances
func MakeEndpoints(s services.Service) Endpoints {
	return Endpoints{
		SendNotification:     makeSendNotificationEndpoint(s),
		SeenNotification:     makeSeenNotificationEndpoint(s),
		GetNotifications:     makeGetNotificationsEndpoint(s),
		SeenAllNotifications: makeSeenAllNotificationsEndpoint(s),
	}
}
