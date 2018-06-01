package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"jyj-go/conf"
	"github.com/gomodule/redigo/redis"
)

/**
 * 定义常量
 */
const (
	WeiboPageSize int = 15
	ArticlePageSize int = 5
)

var DB *gorm.DB
var Redis redis.Conn

/**
 * 初始化连接数据库
 */
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql",conf.GetConfiguration().DataBaseUrl)
	if err != nil {
		return nil, err
	}
	DB = db
	// 其余操作
	DB.SingularTable(true) // 禁用表名负数
	return db, err
}

/**
 * 初始化连接redis
 */
func InitRedis() (redis.Conn, error) {
	c, err := redis.Dial("tcp", conf.GetConfiguration().RedisHost)
	redis.DialPassword(conf.GetConfiguration().RedisPassword)
	if err != nil {
		return nil, err
	}
	Redis = c
	return c, err
}
