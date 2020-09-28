package global

import (
	"library/pkg/logger"
	"library/pkg/setting"

	"github.com/blevesearch/bleve"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	Index           *bleve.Index
)
