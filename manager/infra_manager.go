package manager

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"gokost.com/m/authenticator"
	"gokost.com/m/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type InfraManager interface {
	PostgreConn() *sqlx.DB
	MysqlConn() *gorm.DB
	RedisConn() (context.Context, *redis.Client)
	ConfigToken(tokenConfig authenticator.TokenConfig) authenticator.Token
}

type infraManager struct {
	postgreConn *sqlx.DB
	mysqlConn   *gorm.DB
	redisConn   *redis.Client
	ctx         context.Context
}

func (i *infraManager) PostgreConn() *sqlx.DB {
	return i.postgreConn
}

func (i *infraManager) MysqlConn() *gorm.DB {
	return i.mysqlConn
}

func (i *infraManager) RedisConn() (context.Context, *redis.Client) {
	return i.ctx, i.redisConn
}

func (i *infraManager) ConfigToken(tokenConfig authenticator.TokenConfig) authenticator.Token {
	return authenticator.NewToken(tokenConfig, i.ctx, i.redisConn)
}

func NewInfraManager(configDatabase *config.ConfigDatabase) InfraManager {
	urlPostgresql := configDatabase.PostgreConn()
	urlMysql := configDatabase.MysqlConn()
	redisConfig := configDatabase.RedisConfig()
	conn, err := sqlx.Connect("pgx", urlPostgresql)
	connMysql, errMysql := gorm.Open(mysql.Open(urlMysql), &gorm.Config{})
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})
	ctx := context.Background()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	if errMysql != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", errMysql)
		os.Exit(1)
	}
	return &infraManager{
		postgreConn: conn,
		mysqlConn:   connMysql,
		redisConn:   rdb,
		ctx:         ctx,
	}
}
