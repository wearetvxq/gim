package service
//
//import (
//	"gim/internal/logic/cache"
//	"gim/internal/logic/dao"
//	"gim/internal/logic/model"
//	client "gim/pkg/rpc_cli"
//
//	"gim/pkg/imerror"
//	"gim/pkg/logger"
//	"gim/pkg/pb"
//	"gim/pkg/util"
//
//	"go.uber.org/zap"
//	"context"
//)
//
//// 感觉还是不能这样玩  这都废弃的版本了  修正了看一下 就改回来吧

// 主要1 是传输的解构变了   导致 入参  和 dao层变了  然后再试发送的协议变了  要带出口协议
//
//type httpMessageService struct{}
//
//var HttpMessageService = new(httpMessageService)
//
//// Add 添加消息
//func (httpMessageService) Add(ctx context.Context, message model.Message) error {
//	return dao.MessageDao.Add(ctx, "message", message)
//}
//
//// ListByUserIdAndSeq 查询消息
//func (*httpMessageService) ListByUserIdAndSeq(ctx context.Context, appId, userId, seq int64) ([]model.Message, error) {
//	messages, err := dao.MessageDao.ListBySeq(ctx, "message", appId, model.MessageObjectTypeUser, userId, seq)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return nil, err
//	}
//
//	return messages, nil
//}
//
//// Send 消息发送
//func (s *httpMessageService) Send(ctx context.Context, appId, userId, deviceId int64, send model.SendMessage) error {
//	switch send.ReceiverType {
//	case pb.ReceiverType_RT_USER:
//		err := HttpMessageService.SendToFriend(ctx, appId, userId, deviceId, send)
//		if err != nil {
//			logger.Sugar.Error(err)
//			return err
//		}
//	case pb.ReceiverType_RT_NORMAL_GROUP:
//		err := HttpMessageService.SendToGroup(ctx, appId, userId, deviceId, send)
//		if err != nil {
//			logger.Sugar.Error(err)
//			return err
//		}
//	case pb.ReceiverType_RT_LARGE_GROUP:
//		err := HttpMessageService.SendToChatRoom(ctx, appId, userId, deviceId, send)
//		if err != nil {
//			logger.Sugar.Error(err)
//			return err
//		}
//	}
//
//	return nil
//}
//
//// SendToUser 消息发送至用户
//func (*httpMessageService) SendToFriend(ctx context.Context, appId, userId, deviceId int64, send model.SendMessage) error {
//	// 发给发送者
//	err := HttpMessageService.StoreAndSendToUser(ctx, appId, userId, deviceId, userId, send)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	// 发给接收者
//	err = HttpMessageService.StoreAndSendToUser(ctx, appId, userId, deviceId, send.ReceiverId, send)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	return nil
//}
//
//// SendToGroup 消息发送至群组（使用写扩散）
//func (*httpMessageService) SendToGroup(ctx context.Context, appId, userId, deviceId int64, send model.SendMessage) error {
//	users, err := GroupUserService.GetUsers(ctx, appId, send.ReceiverId)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	if !IsInGroup(users, userId) {
//		logger.Sugar.Error(ctx, appId, send.ReceiverId, "不在群组内")
//		return imerror.ErrNotInGroup
//	}
//
//	// 将消息发送给群组用户，使用写扩散
//	for _, user := range users {
//		err = HttpMessageService.StoreAndSendToUser(ctx, appId, userId, deviceId, user.UserId, send)
//		if err != nil {
//			logger.Sugar.Error(err)
//			return err
//		}
//	}
//	return nil
//}
//
//func IsInGroup(users []model.GroupUser, userId int64) bool {
//	for i := range users {
//		if users[i].UserId == userId {
//			return true
//		}
//	}
//	return false
//}
//
//// SendToChatRoom 消息发送至聊天室（读扩散）
//func (*httpMessageService) SendToChatRoom(ctx context.Context, appId, userId, deviceId int64, send model.SendMessage) error {
//	userIds, err := cache.LargeGroupUserCache.Members(appId, send.ReceiverId)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	isMember, err := cache.LargeGroupUserCache.IsMember(appId, send.ReceiverId, userId)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	if !isMember {
//		logger.Logger.Error("not int group", zap.Int64("app_id", appId), zap.Int64("group_id", send.ReceiverId),
//			zap.Int64("user_id", userId))
//		return imerror.ErrNotInGroup
//	}
//
//	seq, err := SeqService.GetGroupNext(ctx, appId, send.ReceiverId)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	selfMessage := model.Message{
//		MessageId:      send.MessageId,
//		AppId:          appId,
//		ObjectType:     model.MessageObjectTypeGroup,
//		ObjectId:       send.ReceiverId,
//		SenderType:     int32(pb.SenderType_ST_USER),
//		SenderId:       userId,
//		SenderDeviceId: deviceId,
//		ReceiverType:   int32(send.ReceiverType),
//		ReceiverId:     send.ReceiverId,
//		ToUserIds:      model.FormatUserIds(send.ToUserIds),
//		Type:           send.MessageBody.MessageType,
//		Content:        send.MessageBody.MessageContent,
//		Seq:            seq,
//		SendTime:       util.UnunixMilliTime(send.SendTime),
//		Status:         int32(pb.MessageStatus_MS_NORMAL),
//	}
//
//	err = HttpMessageService.Add(ctx, selfMessage)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	messageItem := pb.MessageItem{
//		MessageId:      send.MessageId,
//		SenderType:     pb.SenderType_ST_USER,
//		SenderId:       userId,
//		SenderDeviceId: deviceId,
//		ReceiverType:   send.ReceiverType,
//		ReceiverId:     send.ReceiverId,
//		ToUserIds:      send.ToUserIds,
//		MessageBody:    send.PbBody,
//		Seq:            seq,
//		SendTime:       send.SendTime,
//		Status:         pb.MessageStatus_MS_NORMAL,
//	}
//
//	// 将消息发送给群组用户，使用读扩散
//	for i := range userIds {
//		err = HttpMessageService.SendToUser(ctx, appId, userId, deviceId, userIds[i].UserId, &messageItem)
//		if err != nil {
//			logger.Sugar.Error(err)
//			return err
//		}
//	}
//	return nil
//}
//
//// StoreAndSendToUser 将消息持久化到数据库,并且消息发送至用户
//func (*httpMessageService) StoreAndSendToUser(ctx context.Context, appId, userId, deviceId, toUserId int64, send model.SendMessage) error {
//	logger.Logger.Debug("message_store_send_to_user",
//		zap.String("message_id", send.MessageId),
//		zap.Int64("app_id", appId),
//		zap.Int64("to_user_id", toUserId))
//	seq, err := SeqService.GetUserNext(ctx, appId, userId)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	selfMessage := model.Message{
//		MessageId:      send.MessageId,
//		AppId:          appId,
//		ObjectType:     model.MessageObjectTypeUser,
//		ObjectId:       toUserId,
//		SenderType:     int32(pb.SenderType_ST_USER),
//		SenderId:       userId,
//		SenderDeviceId: deviceId,
//		ReceiverType:   int32(send.ReceiverType),
//		ReceiverId:     send.ReceiverId,
//		ToUserIds:      model.FormatUserIds(send.ToUserIds),
//		Type:           send.MessageBody.MessageType,
//		Content:        send.MessageBody.MessageContent,
//		Seq:            seq,
//		SendTime:       util.UnunixMilliTime(send.SendTime),
//		Status:         int32(pb.MessageStatus_MS_NORMAL),
//	}
//
//	err = HttpMessageService.Add(ctx, selfMessage)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	messageItem := pb.MessageItem{
//		MessageId:      send.MessageId,
//		SenderType:     pb.SenderType_ST_USER,
//		SenderId:       userId,
//		SenderDeviceId: deviceId,
//		ReceiverType:   send.ReceiverType,
//		ReceiverId:     send.ReceiverId,
//		ToUserIds:      send.ToUserIds,
//		MessageBody:    send.PbBody,
//		Seq:            seq,
//		SendTime:       send.SendTime,
//		Status:         pb.MessageStatus_MS_NORMAL,
//	}
//
//	// 查询用户在线设备
//	devices, err := DeviceService.ListOnlineByUserId(ctx, appId, toUserId)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	for i := range devices {
//		// 消息不需要投递给发送消息的设备
//		if deviceId == devices[i].DeviceId {
//			continue
//		}
//
//		message := pb.Message{Message: &messageItem}
//		_, err = client.ConnectRpcClient.Message(devices[i].DeviceId, message)
//		if err != nil {
//			logger.Sugar.Error(err)
//			return err
//		}
//
//		logger.Logger.Debug("message_store_send_to_device",
//			zap.String("message_id", messageItem.MessageId),
//			zap.Int64("app_id", appId),
//			zap.Int64("device_id:", deviceId),
//			zap.Int64("user_id", userId),
//			zap.Int64("seq", messageItem.Seq))
//	}
//	return nil
//}
//
//// SendToUser 消息发送至用户
//func (*httpMessageService) SendToUser(ctx context.Context, appId, userId, deviceId, toUserId int64, messageItem *pb.MessageItem) error {
//	logger.Logger.Info("message_send_to_user",
//		zap.String("message_id", messageItem.MessageId),
//		zap.Int64("app_id", appId),
//		zap.Int64("to_user_id", toUserId),
//		zap.Int64("seq", messageItem.Seq))
//	// 查询用户在线设备
//	devices, err := DeviceService.ListOnlineByUserId(ctx, appId, toUserId)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return err
//	}
//
//	for i := range devices {
//		// 消息不需要投递给发消息的设备
//		if deviceId == devices[i].DeviceId {
//			continue
//		}
//
//		message := pb.Message{Message: messageItem}
//		// TODO 根据设备ID选择连接层服务器
//		_, err = client.ConnectRpcClient.Message(devices[i].DeviceId, message)
//		if err != nil {
//			logger.Sugar.Error(err)
//			return err
//		}
//
//		logger.Logger.Info("message_send_to_device",
//			zap.String("message_id", messageItem.MessageId),
//			zap.Int64("app_id", appId),
//			zap.Int64("device_id:", deviceId),
//			zap.Int64("user_id", userId),
//			zap.Int64("seq", messageItem.Seq))
//	}
//	return nil
//}
//
//func (*httpMessageService) GetMessageSeqs(messages []model.Message) []int64 {
//	seqs := make([]int64, 0, len(messages))
//	for i := range messages {
//		seqs = append(seqs, messages[i].Seq)
//	}
//	return seqs
//}
