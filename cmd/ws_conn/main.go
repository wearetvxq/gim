package main

import (
	"gim/api/ws_conn"
	"gim/config"
	ws_conn2 "gim/internal/ws_conn"
	"gim/pkg/rpc_cli"
	"gim/pkg/util"
	"gim/pkg/logger"
)

func main() {
	// 启动rpc服务
	go func() {
		defer util.RecoverPanic()
		ws_conn.StartRPCServer()  // login启动了两个连接rpc 		ConnRPCAddrs:           "addrs:///127.0.0.1:60000,127.0.0.1:60001",
	}()
	logger.Sugar.Info(1223)
	// 初始化Rpc Client
	rpc_cli.InitLogicIntClient(config.WSConf.LogicRPCAddrs)  //和conn 一样的客户端 连到logic 内部使用 没问题

	// 启动长链接服务器
	ws_conn2.StartWSServer(config.WSConf.WSListenAddr)
}
