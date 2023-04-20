package transport

import (
	"context"
	"notification_service/endpoints"
	"notification_service/pb"
)

func decodeSeenAllNotificationsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SeenAllNotificationsRequest)
	return endpoints.SeenAllNotificationsRequest{
		UserId: req.UserId,
	}, nil
}

func encodeSeenAllNotificationsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.SeenAllNotificationsResponse)
	return &pb.SeenAllNotificationsResponse{
		Success: resp.Success,
	}, nil
}

func (g gRPCServer) SeenAllNotifications(ctx context.Context, request *pb.SeenAllNotificationsRequest) (*pb.SeenAllNotificationsResponse, error) {
	_, resp, err := g.seenAllNotifications.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SeenAllNotificationsResponse), nil
}
