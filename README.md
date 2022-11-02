# common-test
> common包的示例用法，使用github.com/tmnhs/common包快速的搭建一个web服务，

## 1. 使用方法

```shell
#克隆项目
git clone https://github.com/tmnhs/common-test.git

#编译项目
make

#脚本运行项目
#脚本语法：./server.sh {start|stop|restart} {testing|production} 
#默认使用testing配置文件
./server.sh restart 
```

运行成功后可以访问浏览器http://localhost:8089/ping

若得到”pong“,则说明服务启动成功,之后便可以进行**二次开发**了

## 2. 二次开发

### 2.1 入口函数

```go
func main() {
    //参数为需要启动的服务(etcd/mysql/redis) 
    //连接成功后可以通过dbclient.GetMysqlDD(),etcdClient.GetEtcd(),redisclient.GetRedis()获取对应的client
    //通过logger.GetLogger()获取日志处理器
    //通过common.GetConfigModels()获取配置文件的信息
	srv, err := server.NewApiServer(server.WithEtcd(),server.WithMysql(),server.WithRedis())
	if err != nil {
		logger.GetLogger().Error(fmt.Sprintf("new api server error:%s", err.Error()))
		os.Exit(1)
	}
	// 注册路由
	srv.RegisterRouters(handler.RegisterRouters)

	// 建表，当然，如果不需要可以直接注释掉
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

```

### 2.2目录结构

| 目录          | 说明          |
| ----------- | ----------- |
| cmd         | 入口函数        |
| conf        | 配置文件目录      |
| internal    | 业务逻辑目录      |
| haddler     | 路由处理        |
| middlerware | 中间件         |
| model       | 结构体（请求/数据库） |
| service     | 一些业务逻辑服务    |


## 3. 问题
如有问题欢迎加qq:1685290935一起交流讨论