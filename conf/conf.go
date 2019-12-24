package conf

import "os"

// connect和logic公用配置
var (
	MySQL   = "root:Gau43mL9_ff@tcp(localhost:3306)/im2?charset=utf8&parseTime=true"
	NSQIP   = "127.0.0.1:4150"
	RedisIP = "127.0.0.1:6379"

	LogicRPCServerIP   = "127.0.0.1:60000"
	ConnectRPCServerIP = "127.0.0.1:60001"
)

// connect配置
var (
	ConnectTCPListenIP   = "127.0.0.1"
	ConnectTCPListenPort = "50000"
)

// logic配置
var (
	LogicHTTPListenIP = "127.0.0.1:8000"
)

func init() {
	env := os.Getenv("im_env")
	if env == "dev" {
		initDevelopConf()
	}

	if env == "pro" {
		initProductConf()
	}
}

func initDevelopConf() {

}

func initProductConf() {

}
