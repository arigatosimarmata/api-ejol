package config

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := string(os.Getenv("DB_USER"))
	dbPass := string(os.Getenv("DB_PASSWORD"))
	dbHost := string(os.Getenv("DB_HOST"))
	dbPort := string(os.Getenv("DB_PORT"))
	dbName := string(os.Getenv("DB_NAME"))

	dbTimeout := string(os.Getenv("DB_TIMEOUT"))
	dbMaxConn, err := strconv.Atoi(os.Getenv("DB_MAX_CONN"))
	if err != nil {
		return nil, err
	}

	dbMaxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		return nil, err
	}

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?timeout="+dbTimeout+"s")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.SetMaxOpenConns(dbMaxConn)
	db.SetMaxIdleConns(dbMaxIdleConn)

	return db, nil
}
