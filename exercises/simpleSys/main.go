package main

import (
	"database/sql"
	"fmt"
	"learning-go/exercises/simpleSys/stuinfo"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //声明一个全局的 db 变量

// 初始化 MySQL 函数
func initMySQL() (err error) {
	dsn := "root:lianhao816@tcp(127.0.0.1:3306)/CollegeManagement"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 查询一行数据
func queryRowDemo() (u1 *stuinfo.Student, err error) {
	// 声明查询语句
	sqlStr := "SELECT StudentID,Name,Major FROM Student WHERE Major = ?"
	// 声明一个 user 类型的变量
	var u stuinfo.Student
	// 执行查询并且扫描至 u
	err = db.QueryRow(sqlStr, "软件工程").Scan(&u.StudentID, &u.Name, &u.Major)
	if err != nil {
		return nil, err
	}
	u1 = &u
	return
}

func main() {
	// 初始化 MySQL
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	u1, err := queryRowDemo()
	if err != nil {
		fmt.Printf("err:%s", err)
	}
	fmt.Printf("学号:%s, 姓名:%s, 专业:%s\n", u1.StudentID, u1.Name, u1.Major)
}
