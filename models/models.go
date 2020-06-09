package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// 私密
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/leeexing/go-social/pkg/logging"
	"github.com/leeexing/go-social/pkg/setting"
)

var db *gorm.DB

// Model 模型结构体
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `gorm:"create_on"`
	ModifiedOn int `gorm:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		logging.Fatal(2, "Fail to get section 'database': %v", sec)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		logging.Info(err)
	}

	// gorm 默认表名的一个处理函数
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true) // 禁用复数. 就是表名后面不添加 s
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB 关闭数据库连接
func CloseDB() {
	defer db.Close()
}
