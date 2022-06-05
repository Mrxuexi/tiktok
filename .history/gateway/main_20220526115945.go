package main

import (
	"fmt"
	"os"
	"tiktok/gateway/controller"
	"tiktok/gateway/routers"
	"tiktok/pkg/logger"
	"tiktok/pkg/mymysql"
	"tiktok/pkg/myredis"
	"tiktok/setting"
	"tiktok/pkg/snowflake"

	"go.uber.org/zap"
)

func initApp() error {
	if err := setting.Init(os.Args[1]); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		return err
	}

	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		fmt.Printf("init logger failed, err:%v\n", err)
		return err
	}

	if err := mymysql.InitMysql(setting.Conf.MySQLConfig); err != nil {
		zap.L().Fatal("load config failed, err:", zap.Error(err))
		return err
	}

	if err := myredis.Init(setting.Conf.RedisConfig); err != nil {
		zap.L().Fatal("init redis failed, err:", zap.Error(err))
		return err
	}
	defer myredis.Close()

	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		zap.L().Fatal("init redis failed, err:", zap.Error(err))
		return err
	}


	// 注册路由
	r := routers.SetupRouter(setting.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return err
	}
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: bluebell config.yaml")
		return
	}

}