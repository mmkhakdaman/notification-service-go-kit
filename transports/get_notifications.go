package transport

import (
	"context"
	"notification_service/endpoints"
	"notification_service/pb"
)

func decodeGetNotificationsRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetNotificationsRequest)
	return endpoints.GetNotificationsRequest{
		UserId: req.UserId,
		Offset: req.Offset,
		Limit:  req.Limit,
	}, nil
}

func encodeGetNotificationsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetNotificationsResponse)
	return &pb.GetNotificationsResponse{
		Success: resp.Success,
	}, nil
}

func (g gRPCServer) GetNotifications(ctx context.Context, request *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	_, resp, err := g.sendNotification.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetNotificationsResponse), nil
}
