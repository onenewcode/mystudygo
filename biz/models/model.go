package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mystudy/pkg/setting"
)

// 数据库初始化
func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName)),
		&gorm.Config{
			//开启sql语句缓存
			PrepareStmt:            false,
			SkipDefaultTransaction: true,
			//NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // 表名前缀，`User`表为`t_users`
			//	SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			//NameReplacer: strings.NewReplacer("CID", "Cid"), // 在转为数据库名称之前，使用NameReplacer更改结构/字段名称
			//},
		})
	if err != nil {
		log.Fatalf("model.Setup err: %v", err)
	}
	sqlDB, err := db.DB()

	// 设置空闲连接池中的最大连接数。
	sqlDB.SetMaxIdleConns(databaseSetting.MaxIdleConns)

	// 设置与数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(databaseSetting.MaxIdleConns)

	//默认操作的表名
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return setting.DatabaseSetting.TablePrefix + defaultTableName
	//}
	return db, nil
}
