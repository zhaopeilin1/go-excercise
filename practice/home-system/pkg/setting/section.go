package setting

import (
	"fmt"
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize     int
	MaxPageSize         int
	LogSavePath         string
	LogFileName         string
	LogFileExt          string
	UploadSavePath      string
	UploadServerUrl     string
	BlevePath           string
	UploadBookAllowExts []string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}
type JWTSettingS struct {
	AppKey string
	Secret string
	Issuer string
	Expire time.Duration
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	fmt.Println(k, v)
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
