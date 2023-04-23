package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammad-firmansyah/jobhun-backend/api"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "jobhun-academy"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	fmt.Print("db connected")
	app := api.SetupRoute(db)

	app.Listen(":5000")
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/jobhun-academy")

	if err != nil {
		return nil, err
	}

	return db, nil
}
