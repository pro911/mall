package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
	"mall/settings"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "f", "", "传入配置文件路径")

	//解析命令行参数
	flag.Parse()
	log.Printf("configFile:%v\n", configFile)

	//1.加载配置文件
	if err := settings.Init(configFile); err != nil {
		log.Printf("init settings failed,err:%v\n", err)
		return
	}

	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "internal/pkg/dal/query",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Initialize a *gorm.DB instance
	dsn := fmt.Sprintf("%v%v", settings.Conf.SQLiteConfig.Host, settings.Conf.SQLiteConfig.DBName)
	log.Printf("dsn: %v\n", dsn)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(
		g.GenerateModel("user"),
		g.GenerateModel("goods"),
		g.GenerateModel("comment"),
		g.GenerateModel("cart"),
		g.GenerateModel("coupon"),
		g.GenerateModel("order"),
	)

	g.Execute()
}
