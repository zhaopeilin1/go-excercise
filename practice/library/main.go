package main

import (
	"library/global"
	"library/internal/model"
	"library/internal/routers"
	"library/pkg/logger"
	"library/pkg/setting"
	"log"
	"net/http"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/blevesearch/bleve"
	// "github.com/blevesearch/bleve/mapping"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}

	err = setupDbEngine()
	if err != nil {
		log.Fatalf("init.setupDbEngine err:%v", err)
	}

	err = setupBleve()
	if err != nil {
		log.Fatalf("init.setupBleve err:%v", err)
	}
}

// @title 博客系统
// @version 1.0
// @description Go编程之旅，一起用go做项目
// @termsOfService https://github.com/go-programing-tour-book
func main() {
	global.Logger.Infof("%s: go-programing-tour-book/%s", "eddycjy", "library")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDbEngine() error {
	var err error
	global.DBEgine, err = model.NewDbEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.DBEgine.AutoMigrate(&model.Tag{}, &model.Article{}, &model.ArticleTag{}, &model.Book{}, &model.Chapter{})
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupBleve() error {
	// mapping := bleve.NewIndexMapping()
	//index, _ := bleve.New(global.AppSetting.BlevePath+"/library.bleve", mapping)
	index, _ := bleve.Open(global.AppSetting.BlevePath + "/library.bleve")
	global.Index = &index
	return nil

}
