package config

import (
	"gim/pkg/logger"

	"go.uber.org/zap"
)

func initLocalConf() {
	LogicConf = logicConf{
		MySQL:                  "root:Gau43mL9_ff@tcp(localhost:3306)/wsgim?charset=utf8&parseTime=true",
		NSQIP:                  "127.0.0.1:4150",  // 又开始用nsq了??
		RedisIP:                "127.0.0.1:6379",
		RPCIntListenAddr:       ":50000",
		ClientRPCExtListenAddr: ":50001",
		ServerRPCExtListenAddr: ":50002",
		ConnRPCAddrs:           "addrs:///127.0.0.1:60000,127.0.0.1:60001",
		LogicHTTPListenIP:      "127.0.0.1:8000",
	}

	ConnConf = connConf{
		TCPListenAddr: ":8080",
		RPCListenAddr: ":60000",
		LocalAddr:     "127.0.0.1:60000",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	WSConf = wsConf{
		WSListenAddr:  ":8081",
		RPCListenAddr: ":60001",
		LocalAddr:     "127.0.0.1:60001",
		LogicRPCAddrs: "addrs:///127.0.0.1:50000",
	}

	logger.Leavel = zap.DebugLevel
	logger.Target = logger.Console
}
