package main

import (
	chat "GoYin/server/kitex_gen/chat/chatservice"
	"GoYin/server/service/chat/config"
	"GoYin/server/service/chat/dao"
	"GoYin/server/service/chat/initialize"
	"GoYin/server/service/chat/pkg"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/utils"
	"github.com/cloudwego/kitex/server"
	"net"
)

func main() {
	initialize.InitLogger()
	r, info := initialize.InitNacos()
	db := initialize.InitDB()
	publisher := initialize.InitProducer()
	subscriber := initialize.InitSubscriber()
	go func() {
		err := pkg.SubscriberManager.Subscribe(*pkg.NewSubscriberManager(subscriber), context.Background(), dao.NewMysqlManager(db))
		if err != nil {
			klog.Error(err)
		}
	}()
	impl := &ChatServiceImpl{
		MysqlManager: dao.NewMysqlManager(db),
		Publisher:    pkg.NewPublishManager(publisher),
		Subscriber:   pkg.NewSubscriberManager(subscriber),
	}
	svr := chat.NewServer(impl,
		server.WithServiceAddr(utils.NewNetAddr("tcp", net.JoinHostPort(config.GlobalServerConfig.Host, config.GlobalServerConfig.Port))),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 2000, MaxQPS: 500}),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GlobalServerConfig.Name}))

	err := svr.Run()

	if err != nil {
		klog.Error(err)
	}
}