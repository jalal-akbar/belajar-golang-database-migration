package app

import (
	"database/sql"
	"time"

	"github.com/jalal-akbar/golang-restful-api/helper"
)

func NewDB() *sql.DB {
	var (
		driverName     = "mysql"
		dataSourceName = "root:root@tcp(localhost:3306)/belajar_golang_database_migration"
	)
	db, err := sql.Open(driverName, dataSourceName)
	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(1 * time.Hour)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	return db
}

// command run database migration
// migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up
// migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force
