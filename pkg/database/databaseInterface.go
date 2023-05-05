package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type DBInf struct {
	db *sqlx.DB
}

var dbIn DBInf

// 연결부
func (d *DBInf) Connect() error {
	dsn := buildDataSourceName()
	if dsn == "" {
		return fmt.Errorf("database connection error: dsn must not be empty")
	}

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	d.db = db
	return nil
}

func buildDataSourceName() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName := os.Getenv("DB_DATABASENAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, databaseName) //"root:1234@tcp(127.0.0.1:3307)/"
}
