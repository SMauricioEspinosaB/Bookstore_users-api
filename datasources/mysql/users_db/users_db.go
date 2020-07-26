package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlusersusername = "root"
	mysqluserspassword = "Meli2020."
	mysqlusershost     = "127.0.0.1:3306"
	mysqlusersschema   = "users_db"
)

var (
	Client   *sql.DB
	username = "root"
	password = "Meli2020."
	host     = "127.0.0.1:3306"
	schema   = "users_db"

	/*
		username = os.Getenv(mysqlusersusername)"
		password = os.Getenv(mysqluserspassword)
		host     = os.Getenv(mysqlusershost)
		schema   = os.Getenv(mysqlusersschema)
	*/
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	if err := Client.Ping(); err != nil {
		log.Fatal(err)
		//panic(err)
	}

	log.Println("database successfully configured")
}
