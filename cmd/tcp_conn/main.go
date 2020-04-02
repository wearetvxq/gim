package main

import (
	"gim/api/tcp_conn"
	"gim/config"
	tcp_conn2 "gim/internal/tcp_conn"
	"gim/pkg/rpc_cli"
	"gim/pkg/util"
)

func main() {
	// 启动rpc服务
	go func() {
		defer util.RecoverPanic()
		//这个就是内部调用的rpc  conn int
		tcp_conn.StartRPCServer()
	}()

	// 初始化Rpc Client   连接 logic int 的rpc
	rpc_cli.InitLogicIntClient(config.ConnConf.LogicRPCAddrs)

	// 启动长链接服务器
	server := tcp_conn2.NewTCPServer(config.ConnConf.TCPListenAddr, 10)  //acceptGoroutineNum 这个是并发控制么?
	server.Start()

	/*
	这个 tcp2 还是处理的这些啊  和tcp1 的区别呢 ?  这个接受 手机客户端?   tcp1 接受 logic 好像是这样  所以1是对内 2是对外
		switch input.Type {
	case pb.PackageType_PT_SIGN_IN:
		h.SignIn(ctx, input)
	case pb.PackageType_PT_SYNC:
		h.Sync(ctx, input)
	case pb.PackageType_PT_HEARTBEAT:
		h.Heartbeat(ctx, input)
	case pb.PackageType_PT_MESSAGE:
		h.MessageACK(ctx, input)
	default:
		logger.Logger.Error("handler switch other")
	}
	return

	*/
}
