package homework011

import (
	"fmt"
	"log"

	"xorm.io/xorm"
)

// 問題5

// User構造体を実装してください。

func insertUser(engine xorm.Engine) {
	user := User{
		ID:    1,
		Name:  "田中太郎",
		Class: 1,
		Age:   20,
	}
	_, err := engine.Table("user").Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("< Insert > ID:%d 名前:%s, Class:%d, 年齢:%d\n", user.ID, user.Name, user.Class, user.Age)
}

// XORMを使って、問題5の仕様を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func selectUser() {
}

func deleteUser(engine xorm.Engine) {
	user := User{}
	result, err := engine.Where("class=?", 1).Delete(&user)
	if err != nil {
		log.Println(err)
	}
	if result == 0 {
		log.Fatal("Not Found")
	}
	fmt.Println("user:", user)
}

func Task5() {
	fmt.Println("homework011_5")

	engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}

	// 各関数を実行します
	insertUser(*engine)

	// selectUser関数を実行してください。

	deleteUser(*engine)
}
