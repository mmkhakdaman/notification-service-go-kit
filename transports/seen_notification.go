package transport

import (
	"context"
	"notification_service/endpoints"
	"notification_service/pb"
)

func decodeSeenNotificationRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SeenNotificationRequest)
	return endpoints.SeenNotificationRequest{
		NotificationId: req.Id,
	}, nil
}

func encodeSeenNotificationResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.SeenNotificationResponse)
	return &pb.SeenNotificationResponse{
		Success: resp.Success,
	}, nil
}

func (g gRPCServer) SeenNotification(ctx context.Context, request *pb.SeenNotificationRequest) (*pb.SeenNotificationResponse, error) {
	_, resp, err := g.sendNotification.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SeenNotificationResponse), nil
}
