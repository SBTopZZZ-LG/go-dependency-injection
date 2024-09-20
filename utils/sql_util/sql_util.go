package sql_util

import (
	"database/sql"
	"fmt"
	"todo_app/config"
)

func CreateConnection(dbConf *config.DBConfig) (*sql.DB, error) {
	sqlDataSourceName := constructSqlDataSourceName(dbConf)

	db, err := sql.Open("mysql", sqlDataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) error {
	return db.Close()
}

func constructSqlDataSourceName(dbConf *config.DBConfig) string {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Name, dbConf.Params)
	return dataSourceName
}
