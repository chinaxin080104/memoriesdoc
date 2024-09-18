package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
	"gozero/api/common/init_db"
	"gozero/api/internal/config"
	"gozero/api/model"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gorm.DB
	UsersModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	//获取mysql链接
	mysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	mysqlDb.AutoMigrate(&model.User{})

	return &ServiceContext{
		Config: c,
		DB:     mysqlDb,
	}
}

func NewServiceContext_back(c config.Config) *ServiceContext {
	//获取mysql链接
	mysqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:     c,
		UsersModel: model.NewUserModel(mysqlConn),
	}
}
