package repository

import (
	"balance/src/main/domain/projection"
	"balance/src/main/infra"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(env *infra.Env) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", env.DatabaseUsername, env.DatabasePassword, env.DatabaseHost, env.DatabasePort, env.DatabaseName)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = database.AutoMigrate(projection.AccountBalanceProjection{})
	if err != nil {
		panic(err)
	}
	return database
}
