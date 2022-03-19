package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

// usersテーブルを作成している
func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.Dbname)

	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)
}

// uuidを生成

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()

	return uuidobj
}

// パスワードをハッシュ化

func Encrpt(plaintext string) (crytext string) {
	crytext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))

	return crytext
}
