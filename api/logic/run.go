package logic

import (
	"gim/config"
	"gim/pkg/pb"
	"gim/pkg/util"
	"net"

	"google.golang.org/grpc"
)


//看错了 这一层在API这里  傻逼了

// StartRpcServer 启动rpc服务
func StartRpcServer() {
	//内部conn 连接调用  只做 消息传递这种业务
	go func() {
		defer util.RecoverPanic()

		intListen, err := net.Listen("tcp", config.LogicConf.RPCIntListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer(grpc.UnaryInterceptor(LogicIntInterceptor))
		pb.RegisterLogicIntServer(intServer, &LogicIntServer{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()

	// 对client 提供的 业务 api
	go func() {
		defer util.RecoverPanic()

		extListen, err := net.Listen("tcp", config.LogicConf.ClientRPCExtListenAddr)
		if err != nil {
			panic(err)
		}
		extServer := grpc.NewServer(grpc.UnaryInterceptor(LogicClientExtInterceptor))
		pb.RegisterLogicClientExtServer(extServer, &LogicClientExtServer{})
		err = extServer.Serve(extListen)
		if err != nil {
			panic(err)
		}
	}()

	//对接入server 提供的业务api
	go func() {
		defer util.RecoverPanic()

		intListen, err := net.Listen("tcp", config.LogicConf.ServerRPCExtListenAddr)
		if err != nil {
			panic(err)
		}
		intServer := grpc.NewServer(grpc.UnaryInterceptor(LogicServerExtInterceptor))
		pb.RegisterLogicServerExtServer(intServer, &LogicServerExtServer{})
		err = intServer.Serve(intListen)
		if err != nil {
			panic(err)
		}
	}()
}
