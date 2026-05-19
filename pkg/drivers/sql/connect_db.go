package sql

import (
	"attendance-api/config"
	"attendance-api/pkg/drivers/common"

	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenDbConnection() *gorm.DB {

	v := common.DbInfo{
		Host:     config.Config("DB_HOST"),
		User:     config.Config("DB_USER"),
		Password: config.Config("DB_PASSWORD"),
		DbName:   config.Config("DB_NAME"),
		Port:     config.Config("DB_PORT"),
	}

	/*pengecekan Host*/
	config := common.NewDbInfo(v)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DbName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed Connection Database : " + err.Error())
	}

	sqlDb, err := db.DB()

	if err != nil {
		panic("Failed to get database object : " + err.Error())
	}

	if err := sqlDb.Ping(); err != nil {
		panic("Database not reachable : " + err.Error())
	}

	return db
}
