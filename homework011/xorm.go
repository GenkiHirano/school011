package homework011

import (
	"fmt"
	"log"

	"xorm.io/xorm"
)

// 問題5
type Student struct {
}

// XORMを使って、問題5の仕様を実装してください。
func InsertSum() {
}

func Delete(engine xorm.Engine) {
	student := Student{}
	result, err := engine.Where("class=?", 1).Delete(&student)
	if err != nil {
		log.Println(err)
	}
	if result == 0 {
		log.Fatal("Not Found")
	}
	fmt.Println("user:", student)
}

func Task5() {
	fmt.Println("homework011_5")

	engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(Student))
	if err != nil {
		log.Fatal(err)
	}

	studentA := Student{
		ID:    1,
		Name:  "田中太郎",
		Class: 1,
		Age:   20,
		Score: 70,
	}

	studentB := Student{
		ID:    2,
		Name:  "鈴木次郎",
		Class: 1,
		Age:   25,
		Score: 80,
	}

	studentC := Student{
		ID:    3,
		Name:  "伊藤桃子",
		Class: 1,
		Age:   30,
		Score: 90,
	}

	students := []Student{studentA, studentB, studentC}

	// InsertSum関数を実行してください。
	Delete(*engine)
}
