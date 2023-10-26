package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/thiagoalvesp/go-todo-app/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disabled",
		conf.Host, conf.Port, conf.User, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	err = conn.Ping()

	return conn, err

}
