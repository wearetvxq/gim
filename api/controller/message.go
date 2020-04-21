package controller

import (
	"gim/internal/logic/model"
	"gim/pkg/imerror"
	"gim/pkg/pb"
	"gim/pkg/util"
	"io/ioutil"
	"net/http"

	"github.com/json-iterator/go"
	"gim/internal/logic/service"
)

func init() {
	g := Engine.Group("/message")
	g.POST("/send", handler(MessageController{}.Send))
}

type MessageController struct{}

//其他的service层 都是传 model 的结构体 所以 这边 message 也确实改了

// Send 发送消息
func (MessageController) Send(c *context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusOK, NewWithError(imerror.WrapErrorWithData(imerror.ErrBadRequest, err)))
		return
	}

	var str interface{}
	jsoniter.Get(bytes, "message_body", "message_content").ToVal(&str)
	body, err := jsoniter.Marshal(str)
	if err != nil {
		c.JSON(http.StatusOK, NewWithError(imerror.WrapErrorWithData(imerror.ErrBadRequest, err)))
		return
	}

	var send model.SendMessage
	err = jsoniter.Unmarshal(bytes, &send)
	if err != nil {
		c.JSON(http.StatusOK, NewWithError(imerror.WrapErrorWithData(imerror.ErrBadRequest, err)))
		return
	}
	send.MessageBody.MessageContent = util.Bytes2str(body)

	pbMessageType := pb.MessageType(send.MessageBody.MessageType)
	if pbMessageType == pb.MessageType_MT_UNKNOWN {
		c.JSON(http.StatusOK, NewWithError(imerror.WrapErrorWithData(imerror.ErrBadRequest, err)))
		return
	}
	send.PbBody = model.NewMessageBody(send.MessageBody.MessageType, send.MessageBody.MessageContent)


	sender:=model.Sender{
		AppId:      c.appId,
		SenderType: pb.SenderType_ST_USER,
		SenderId:   c.userId,
		DeviceId:   c.deviceId,
	}



	//  http的 SendMessage   到 pb
	//type SendMessage struct {
	//	ReceiverType pb.ReceiverType `json:"receiver_type"`
	//	ReceiverId   int64           `json:"receiver_id"`
	//	ToUserIds    []int64         `json:"to_user_ids"`
	//	MessageId    string          `json:"message_id"`
	//	SendTime     int64           `json:"send_time"`
	//	MessageBody  struct {
	//		MessageType    int    `json:"message_type"`
	//		MessageContent string `json:"-"`
	//	} `json:"message_body"`
	//	PbBody *pb.MessageBody `json:"-"`
	//}
	/*真实 所以这里要写不同的分之了??
				MessageBody: &pb.MessageBody{
				MessageType: pb.MessageType_MT_TEXT,
				MessageContent: &pb.MessageContent{
					Content: &pb.MessageContent_Text{
						Text: &pb.Text{
							Text: "test",
						},
					},
				},
			},
	*/
	// 这个协议都改了 不支持json 了 所以这里json 转换就特别麻烦
	in :=pb.SendMessageReq{
		MessageId:            send.MessageId,
		ReceiverType:         send.ReceiverType,
		ReceiverId:           send.ReceiverId,
		ToUserIds:            send.ToUserIds,
		MessageBody:          &pb.MessageBody{
			MessageType:          	pb.MessageType(int32(send.MessageBody.MessageType)),
			MessageContent:       &pb.MessageContent{
				Content: &pb.MessageContent_Text{
					Text: &pb.Text{
						Text: send.MessageBody.MessageContent,
					},
				},
			},
		},
		SendTime:             0,
		IsPersist:            false,
	}
	//c.response(nil, service.HttpMessageService.Send(c, c.appId, c.userId, c.deviceId, send))

	c.response(nil, service.MessageService.Send(c, sender, in))

	// 这个发送 message 的格式改了
}
