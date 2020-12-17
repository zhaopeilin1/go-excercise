package global

import (
	"home-system/pkg/logger"
	"home-system/pkg/setting"

	"github.com/blevesearch/bleve"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	Index           *bleve.Index
	JWTSetting      *setting.JWTSettingS
)
