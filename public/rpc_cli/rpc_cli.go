package rpc_cli

import (
	"context"
	"fmt"
	"gim/conf"
	"gim/public/grpclib"
	"gim/public/logger"
	"gim/public/pb"
	"google.golang.org/grpc"
)

var (
	LogicIntClient   pb.LogicIntClient
	ConnectIntClient pb.ConnIntClient
)

func InitLogicIntClient() {
	logger.Logger.Info("实例化conn的logic cli  50000")
	conn, err := grpc.DialContext(context.TODO(), conf.LogicRPCAddrs, grpc.WithInsecure())
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	LogicIntClient = pb.NewLogicIntClient(conn)
}

func InitConnIntClient() {
	logger.Logger.Info("实例化conn的conn cli 60000")

	conn, err := grpc.DialContext(context.TODO(), conf.ConnRPCAddrs, grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, grpclib.Name)))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}

	ConnectIntClient = pb.NewConnIntClient(conn)
}
