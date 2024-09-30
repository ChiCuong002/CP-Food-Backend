package initialize

import (
	"database/sql"
	"fmt"
	"food-recipes-backend/global"

	"go.uber.org/zap"
	_ "github.com/lib/pq"
)

func InitDatabase() {
	// Define the connection string
	m := global.Config.PostgreSQL
	dsn := "user=%s dbname=%s sslmode=disable password=%s host=%s port=%s"
	connStr := fmt.Sprintf(dsn, m.User, m.DbName, m.Password, m.Host, m.Port)
	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	checkErrorPanic(err, "error opening database connection")
	// Ping the database to ensure a successful connection
	err = db.Ping()
	checkErrorPanic(err, "error pinging database")
	global.Logger.Info("Successfully connected to the database")
	global.Db = db
}

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err)) //write into log
		panic(err)
	}
}