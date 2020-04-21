package main

import (
	"gim/api/logic"
	"gim/api/controller"
	"gim/config"
	"gim/internal/logic/db"
	"gim/pkg/logger"
	"gim/pkg/rpc_cli"
	"gim/pkg/util"
)

func main() {
	// 初始化数据库
	db.InitDB()

	// 初始化自增id配置
	util.InitUID(db.DBCli)

	// 初始化RpcClient
	logger.Logger.Info(config.LogicConf.ConnRPCAddrs)
	rpc_cli.InitConnIntClient(config.LogicConf.ConnRPCAddrs)

	logic.StartRpcServer()
	logger.Logger.Info("logic server start")

	// 启动web容器
	err := controller.Engine.Run(config.LogicConf.LogicHTTPListenIP)
	if err != nil {
		logger.Sugar.Error(err)
	}


	select {}
}
