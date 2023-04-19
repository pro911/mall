package initializ

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"mall/settings"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Redis 初始化连接
func Redis(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DBName,
		PoolSize: cfg.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err = rdb.Ping(ctx).Result()
	return
}

// RedisDB 外部redis访问
func RedisDB() *redis.Client {
	return rdb
}

// RedisClose 关闭客户端，释放所有打开的资源。
// 关闭客户端很少见，因为客户端是长期存在的，
// 并且在许多goroutine之间共享
func RedisClose() {
	_ = rdb.Close()
}
