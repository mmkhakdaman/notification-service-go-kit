package transport

import (
	"context"
	"notification_service/endpoints"
	"notification_service/pb"
)

func decodeSendNotificationRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SendNotificationRequest)
	return endpoints.SendNotificationRequest{
		Title:     req.Title,
		Message:   req.Message,
		Recipient: req.Recipient,
		Sender:    req.Sender,
		Type:      req.Type,
		Avatar:    req.Avatar,
		Link:      req.Link,
		LinkText:  req.LinkText,
		Icon:      req.Icon,
	}, nil
}

func encodeSendNotificationResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.SendNotificationResponse)
	return &pb.SendNotificationResponse{
		Success: resp.Success,
	}, nil
}

func (g gRPCServer) SendNotification(ctx context.Context, request *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
	_, resp, err := g.sendNotification.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SendNotificationResponse), nil
}
