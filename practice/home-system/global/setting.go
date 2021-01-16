package global

import (
	"home-system/pkg/logger"
	"home-system/pkg/setting"

	"github.com/blevesearch/bleve"
	"github.com/go-co-op/gocron"
	"github.com/robfig/cron"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	Index           *bleve.Index
	JWTSetting      *setting.JWTSettingS
	Cron            *cron.Cron
	Scheduler       *gocron.Scheduler
)
