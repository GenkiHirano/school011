package homework011

import (
	"fmt"
	"log"

	"xorm.io/xorm"
)

// 問題6

// Student構造体を実装してください。

// insertSumStudent を実装してください。
// 中身を適宜変更してください。引数も変更して大丈夫です。
func insertSumStudent(engine xorm.Engine, students []Student) {
}

func deleteStudent(engine xorm.Engine) {
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

func Task6() {
	fmt.Println("homework011_6")

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

	// 各関数を実行します

	// InsertSum関数を実行してください。

	deleteStudent(*engine)
}
