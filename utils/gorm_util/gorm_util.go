package gorm_util

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSilentGormInstanceWithMySQLDriver(sqlConn *sql.DB) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		Conn: sqlConn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
