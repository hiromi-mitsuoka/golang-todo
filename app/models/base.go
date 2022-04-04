package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	// モジュールモードでの内部パッケージのimportの記述(https://qiita.com/fetaro/items/31b02b940ce9ec579baf#%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB%E3%83%A2%E3%83%BC%E3%83%89%E3%81%A7%E3%81%AE%E5%86%85%E9%83%A8%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%81%AEimport)
	"github.com/hiromi-mitsuoka/golang-todo/config"

	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
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

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// Change password from plaintext to hash value
func Encrypt(plaintext string) (cryptext string) {
	// Sprintf: You can embed a variable in a string, which returns a new string
	// %x: 16進数
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
