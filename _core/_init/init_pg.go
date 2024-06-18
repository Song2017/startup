package initial

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type PgConn struct {
	DSN           string `mapstructure:"dsn" json:"dsn" yaml:"dsn"`                                  // 是否开启Gorm全局日志
	MaxIdleConns  int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns  int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	SlowThreshold int    `mapstructure:"slow-threshold" json:"slow-threshold" yaml:"slow-threshold"` // slow sql threshold
}

// GormPgSql 初始化 Postgresql 数据库
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func InitGormPg() *gorm.DB {
	p := LoadServerConfig(nil).PgConn
	return p.GormPgSqlByConfig()
}

// GormPgSqlByConfig 初始化 Postgresql 数据库 通过参数
func (p PgConn) GormPgSqlByConfig() *gorm.DB {
	if p.DSN == "" {
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  p.DSN, // DSN data source name
		PreferSimpleProtocol: false,
	}
	optsConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "Gorm",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Duration(p.SlowThreshold) * time.Millisecond,
				LogLevel:      logger.Warn,
				// IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful: true,
			},
		),
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig), &optsConfig)
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	return db
}
