package initializ

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"mall/internal/app/model"
	"mall/settings"
	"time"
)

var sdb *gorm.DB

func SQLite(cfg *settings.SQLiteConfig) (err error) {
	dsn := fmt.Sprintf("%v%v", cfg.Host, cfg.DBName)
	fmt.Println("dsn", dsn)
	sdb, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //为了确保数据一致性，GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升。
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   //表名前缀，`User` 的表名应该是`d_users`
			SingularTable: true, //使用单数表名，启用该选项，此时，`User` 的表名应该是`d_user`
			NameReplacer:  nil,
			NoLowerCase:   false,
		},
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: false, //执行任何 SQL 时都创建 prepared statement 并缓存，可以提高后续的调用速度
	})
	if err != nil {
		log.Fatal(2, "Fail to sqlite.open err:%v\n", err)
	}

	sqliteDB, err := sdb.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqliteDB.SetMaxIdleConns(cfg.MaxIdleConns) //最大空闲连接数
	sqliteDB.SetMaxOpenConns(cfg.MaxOpenConns) //最大连接数
	sqliteDB.SetConnMaxLifetime(time.Hour)     //设置了连接可复用的最大时间

	err = sdb.AutoMigrate(
		&model.User{},
		&model.Goods{},
		&model.Cart{},
		&model.Order{},
		&model.Comment{},
		&model.Coupon{},
	)

	if err != nil {
		log.Fatal(2, "Fail to sqlite.auto_migrate err:%v\n", err)
	}

	//初始化项目数据
	initProject(sdb)
	return
}

// SQLiteDB 返回链接数据库
func SQLiteDB() *gorm.DB {
	return sdb
}

// 初始化迁移一些基础数据
func initProject(sdb *gorm.DB) {
	var count int64
	sdb.Model(&model.User{}).Count(&count)
	if count > 0 {
		return
	}

	users := []*model.User{
		{
			Username: "proxx01@qq.com",
			Password: "123456",
		},
		{
			Username: "proxx02@qq.com",
			Password: "123456",
		},
		{
			Username: "proxx03@qq.com",
			Password: "123456",
		},
	}

	coupons := []*model.Coupon{
		{
			Name:    "满5元减1元",
			Satisfy: 500,
			Minus:   100,
			Desc:    "满5元减1元,限时大放送！",
		},
		{
			Name:    "满100元减10元",
			Satisfy: 10000,
			Minus:   1000,
			Desc:    "满100元减10元,限时大放送！",
		},
	}

	comments := []*model.Comment{
		{
			GoodsID: 1,
			UserID:  1,
			Desc:    "商品很好用!",
		},
		{
			GoodsID: 1,
			UserID:  2,
			Desc:    "商品很好用!满100元减10元",
		},
		{
			GoodsID: 1,
			UserID:  1,
			Desc:    "商品很好用!满100元减10元,限时大放送！",
		},
		{
			GoodsID: 1,
			UserID:  1,
			Desc:    "商品很好用!满100元减10元,限时大放送！",
		},
		{
			GoodsID: 1,
			UserID:  1,
			Desc:    "商品很好用!满100元减10元,限时大放送！",
		},
		{
			GoodsID: 2,
			UserID:  2,
			Desc:    "商品很好用!满100元减10元,限时大放送！",
		},
	}

	goodsS := []*model.Goods{
		{
			Name:    "测试商品01",
			Desc:    "测试商品01 描述",
			Price:   100,
			Details: "测试商品01 商品很好用!满100元减10元,限时大放送！商品很好用!满100元减10元,限时大放送！商品很好用!满100元减10元,限时大放送！",
		},
		{
			Name:    "测试商品02",
			Desc:    "测试商品02 描述",
			Price:   2000,
			Details: "测试商品02 商品很好用!满100元减10元,限时大放送！商品很好用!满100元减10元,限时大放送！商品很好用!满100元减10元,限时大放送！",
		},
	}
	//否则执行初始化数据
	sdb.Model(&model.User{}).Create(users)
	sdb.Model(&model.Coupon{}).Create(coupons)
	sdb.Model(&model.Comment{}).Create(comments)
	sdb.Model(&model.Goods{}).Create(goodsS)
	return
}
