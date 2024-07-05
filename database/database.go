package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DBdir string
	DB    *sql.DB
)

func Init() {
	DBdir = "./database/database/"

	DB, _ = sql.Open("sqlite3", DBdir+"database.db")

	scheme, _ := os.ReadFile(DBdir + "scheme.sql")

	_, _ = DB.Exec(string(scheme))
}

func IsUserExist(email string) (exist bool, id int) {
	var count int
	row := DB.QueryRow("SELECT COUNT(*), id FROM user WHERE email = ?", email)
	row.Scan(&count, &id)

	return count > 0, id
}

func IsUserIdExist(id int) bool {
	var count int
	row := DB.QueryRow("SELECT COUNT(*) FROM user WHERE id = ?", id)
	row.Scan(&count)

	return count > 0
}

func InsertUser(id int, name string, email string, password string) int64 {
	stmt, _ := DB.Prepare("INSERT INTO user (id, name, email, password) VALUES (?, ?, ?, ?)")
	res, err := stmt.Exec(id, name, email, password)
	if err != nil {
		panic(err)
	}
	lastId, _ := res.LastInsertId()

	return lastId
}

func Close() {
	DB.Close()
}
