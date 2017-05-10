package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/yaozijian/MiningOpt/distribution"
)

func runTaskWorker(webcfg *WebConfig) {

	rpcxcfg := distribution.RpcxServerConfig{
		ServiceAddr: fmt.Sprintf("%v:%v", webcfg.MyIpAddr, webcfg.RpcxPort),
		EtcdServers: webcfg.EtcdServers,
		URLPrefix:   fmt.Sprintf("http://%v:%v", webcfg.MyIpAddr, beego.BConfig.Listen.HTTPPort),
	}

	distribution.StartWorker(rpcxcfg)
}
