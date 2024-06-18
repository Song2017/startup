package initial

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Username string `mapstructure:"username" json:"username" yaml:"username"` // 密码
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
}

func InitRedis() *redis.Client {
	redisCfg := LoadServerConfig(nil).Redis
	return redisCfg.InitRedisByConfig()
}
func (redisCfg Redis) InitRedisByConfig() *redis.Client {
	if redisCfg.Addr == "" {
		return nil
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		Username: redisCfg.Username, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("redis connect ping failed, err:", err)
	}
	return client
}
