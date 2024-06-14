package main

import (
	gateway "gateway/kitex_gen/gateway/bizservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {
	bizService := initializeBizService()
	registry, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	serverAddress := resolveServerAddress()

	svr := gateway.NewServer(
		bizService,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "gateway"}),
		server.WithRegistry(registry),
		server.WithServiceAddr(serverAddress),
	)

	if err := svr.Run(); err != nil {
		log.Println(err.Error())
	}
}

func initializeBizService() *BizServiceImpl {
	bizService := new(BizServiceImpl)
	bizService.InitDB()
	return bizService
}

func resolveServerAddress() *net.TCPAddr {
	serverAddress, _ := net.ResolveTCPAddr("tcp", ":9999")
	return serverAddress
}
