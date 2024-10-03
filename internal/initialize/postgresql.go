package initialize

import (
	"fmt"
	"food-recipes-backend/global"

	"database/sql"
	"go.uber.org/zap"
	_ "github.com/lib/pq"	
)

func InitDatabase() {
	// Define the connection string
	m := global.Config.PostgreSQL
	url := "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	connStr := fmt.Sprintf(url, m.User, m.Password, m.Host, m.Port, m.DbName)
    db, err := sql.Open("postgres", connStr)
    checkErrorPanic(err, "error creating database connection pool")
	global.Logger.Info("Successfully connected to the database")
	global.Db = db
}

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err)) //write into log
		panic(err)
	}
}