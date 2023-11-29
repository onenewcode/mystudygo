package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"log"
	"mystudy/biz/models"
	v1 "mystudy/biz/router/v1"
	"mystudy/global"
	"mystudy/pkg/setting"
)

func main() {
	h := server.New(
		server.WithHostPorts(global.ServerSetting.HttpPort),
		server.WithReadTimeout(global.ServerSetting.ReadTimeout),
		server.WithWriteTimeout(global.ServerSetting.WriteTimeout),
	)
	v1.InitRouter(h)
	//h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
	//	ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	//})
	h.Spin()
}
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	//err = setupDBEngine()
	//if err != nil {
	//	log.Fatalf("init.setupDBEngine err: %v", err)
	//}
	//// es初始化
	//cfg := elasticsearch.Config{
	//	Addresses: []string{
	//		"http://192.168.218.134:9200",
	//	},
	//}
	//es, err := elasticsearch.NewClient(cfg)
	//log.Println(err)
	//if err == nil {
	//	log.Println(elasticsearch.Version)
	//	log.Println(es.Info())
	//	global.ESClient=es
	//} else {
	//	log.Println("Something wrong with connection to Elasticsearch")
	//}
}

func setupSetting() error {
	msetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	//err = msetting.ReadSection("Server", &global.ServerSetting)
	err = msetting.ReadSection(
		"Server",
		&global.ServerSetting,
	)
	if err != nil {
		return err
	}
	err = msetting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = msetting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	//global.ServerSetting.ReadTimeout *= time.Second
	//global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = models.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
