package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"models/series"
	"os"
)

var (
	Dbase      *sql.DB
	SeriesRepo series.SeriesRepository
)

func init() {
	dbUser := "ares"
	dbPass := "password"
	dbProto := "tcp"
	dbHost := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	dbPort := "3306"
	dbName := "ares"
	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", dbUser, dbPass, dbProto, dbHost, dbPort, dbName)
	fmt.Println("Opening database with DSN:", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	Dbase = db
	//TODO: Error Checking
	SeriesRepo, _ = series.NewMysqlSeriesRepository(Dbase)
}
