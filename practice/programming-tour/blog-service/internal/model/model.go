package model

import (
	"blog-service/global"
	"blog-service/pkg/setting"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Model struct {
	// id
	Id uint32 `json:id`
	// 创建时间
	CreatedOn uint32 `json:created_on`
	// 创建人
	CreatedBy string `json:created_by`
	// 修改时间
	ModifiedOn uint32 `json:modified_on`
	// 修改人
	ModifiedBy string `json:modified_by`
	// 删除时间
	DeletedOn uint32 `json:deleted_on`
	// 是否删除 0为未删除、1为已删除
	IsDel uint8 `json:is_del`
}

func NewDbEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if databaseSetting.DBType == "sqlite3" {
		db, err = gorm.Open(databaseSetting.DBType, databaseSetting.DBName)
	} else {
		db, err = gorm.Open(databaseSetting.DBType,
			fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
				databaseSetting.UserName,
				databaseSetting.Password,
				databaseSetting.Host,
				databaseSetting.DBName,
				databaseSetting.Charset,
				databaseSetting.ParseTime,
			))
	}

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallBack)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallBack)
	db.Callback().Delete().Replace("gorm:delete", deleteCallBack)

	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db, nil
}

func updateTimeStampForCreateCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallBack(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}

}

func deleteCallBack(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedOnfield, hasDeletdeOnField := scope.FieldByName("DeleteOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")

		if !scope.Search.Unscoped && hasDeletdeOnField && hasIsDelField {
			now := time.Now().Unix()
			scope.Raw(fmt.Sprintf(
				"update %v set %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnfield.DBName),
				scope.AddToVars(now),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"delete from %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}

	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
