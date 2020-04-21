package logic

import (
	"context"
	"gim/internal/logic/model"
	"gim/internal/logic/service"
	"gim/pkg/grpclib"
	"gim/pkg/pb"
)

type LogicServerExtServer struct{}


// 业务服务的话 就只需要有 appid 的请求头

// SendMessage 发送消息
func (*LogicServerExtServer) SendMessage(ctx context.Context, in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	appId, err := grpclib.GetCtxAppId(ctx)
	if err != nil {
		return nil, err
	}

	sender := model.Sender{
		AppId:      appId,
		SenderType: pb.SenderType_ST_BUSINESS,
	}
	err = service.MessageService.Send(ctx, sender, *in)
	if err != nil {
		return nil, err
	}
	return &pb.SendMessageResp{}, nil
}
