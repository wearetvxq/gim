package rpc_cli

import (
	"context"
	"fmt"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/pb"

	"google.golang.org/grpc"
)

var (
	LogicIntClient   pb.LogicIntClient
	ConnectIntClient pb.ConnIntClient
)

func InitLogicIntClient(addr string) {
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	LogicIntClient = pb.NewLogicIntClient(conn)
}


// 只会往 conn 发一个DeliverMessage 嘛
func InitConnIntClient(addr string) {
	// 这里是有一个LoadBalancingPolicy 策略啊  有点牛逼啊 需要看一下
	conn, err := grpc.DialContext(context.TODO(), addr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, grpclib.Name)))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	ConnectIntClient = pb.NewConnIntClient(conn)
}
