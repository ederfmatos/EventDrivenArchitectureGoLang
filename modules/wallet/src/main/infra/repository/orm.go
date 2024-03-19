package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wallet/src/main/infra"
	"wallet/src/main/infra/repository/orm"
)

func GetDatabase(env *infra.Env) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", env.DatabaseUsername, env.DatabasePassword, env.DatabaseHost, env.DatabasePort, env.DatabaseName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(orm.CustomerORM{})
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(orm.AccountORM{})
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(orm.TransactionORM{})
	if err != nil {
		panic(err)
	}
	return database
}
