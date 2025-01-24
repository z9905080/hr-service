package infra

import (
	"fmt"
	"github.com/z9905080/hr_service/environment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func NewGormDB(config environment.Config) (*gorm.DB, error) {

	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(getDialector(config), gormCfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDialector(cfg environment.Config) gorm.Dialector {
	// dsn => user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true&loc=UTC&timeout=%s&readTimeout=%s&writeTimeout=%s",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Database, cfg.DB.ConnectTimeout, cfg.DB.ReadTimeout, cfg.DB.WriteTimeout)

	fmt.Println(dsn)

	return mysql.Open(dsn)
}
