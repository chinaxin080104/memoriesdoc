package svc

import (
	"gorm.io/gorm"
	"gozero/mysqluser/common/init_db"
	"gozero/mysqluser/internal/config"
	"gozero/mysqluser/models"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	err := mysqlDb.AutoMigrate(&models.User{})
	if err != nil {
		return nil
	}
	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
	}
}
