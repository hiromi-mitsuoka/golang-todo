package main

import (
	"fmt"

	// モジュールモードでの内部パッケージのimportの記述(https://qiita.com/fetaro/items/31b02b940ce9ec579baf#%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB%E3%83%A2%E3%83%BC%E3%83%89%E3%81%A7%E3%81%AE%E5%86%85%E9%83%A8%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E3%81%AEimport)
	"github.com/hiromi-mitsuoka/golang-todo/app/models"
)

func main() {
	// Confirmation of ini operation
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test"
	u.Email = "test@example.com"
	u.Password = "testpassword"
	fmt.Println(u)

	u.CreateUser()
}
