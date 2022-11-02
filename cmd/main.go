package main

import (
	"fmt"
	"github.com/tmnhs/common-test/internal/handler"
	"github.com/tmnhs/common-test/internal/service"
	"github.com/tmnhs/common/dbclient"
	"github.com/tmnhs/common/logger"
	"github.com/tmnhs/common/server"
	"os"
)

func main() {
	//是否配置etcd,mysql和redis
	srv, err := server.NewApiServer(server.WithEtcd(),server.WithMysql(),server.WithRedis())
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("new api server error:%s", err.Error()))
		os.Exit(1)
	}
	// Register the API routing service
	srv.RegisterRouters(handler.RegisterRouters)

	//init db table
	err = service.RegisterTables(dbclient.GetMysqlDB())
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("init db table error:%#v", err))
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("startup api server error:%v", err.Error()))
		os.Exit(1)
	}
	os.Exit(0)
}
