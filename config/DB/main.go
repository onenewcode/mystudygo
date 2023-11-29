package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		//  设置输出路径
		OutPath: "./dao",
		//Mode:    gen.WithDefaultQuery, // 选择生成模式
		//选择生成测试代码
		//WithUnitTest: true,
	})
	//  建立数据库连接
	gormdb, _ := gorm.Open(mysql.Open("root:root@(192.168.218.134:3306)/hotel?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // 选择数据库连接
	g.ApplyBasic(
		g.GenerateModelAs("tb_hotel", "Hotel"),
	)
	// 生成代码
	g.Execute()

}
