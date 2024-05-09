package base

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/DB_TEST_1")
	db.Ping()
	defer db.Close()
	if err != nil {
		fmt.Println("数据库连接失败！")
		log.Fatalln(err)
	}
	// 建表
	//createTable(db)
	// 插入数据
	//insertData(db)
	// 更新数据
	//changeData(db)
	// 删除数据
	delData(db)
	// 查询数据
	selectData(db)
}
func createTable(db *sql.DB) {
	query := "CREATE TABLE blog(id INT NOT NULL , title VARCHAR(2000), PRIMARY KEY(ID));"
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("Successfully Created blog\n")
}
func insertData(db *sql.DB) {
	query := "INSERT INTO user VALUES (1,'jelly')"
	_, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Successfully Insert Data TO User\n")
}
func selectData(db *sql.DB) {
	//result, err := db.Query("SELECT * FROM user")
	// 条件查询
	query := "SELECT * FROM user"
	//query := "SELECT * FROM user WHERE id = ?"
	//selectId := 1
	//result, err := db.Query(query, selectId)
	result, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	for result.Next() {
		var id int
		var name string
		err = result.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Id: %d, Name: %s\n", id, name)
	}
}

func changeData(db *sql.DB) {
	query := "UPDATE user SET name = ? WHERE id = ?"
	res, err1 := db.Exec(query, "huguodong", 1)
	if err1 != nil {
		panic(err1)
	}
	affectedRows, err2 := res.RowsAffected()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Update Success, The statement affected %d rows\n", affectedRows)
}
func delData(db *sql.DB) {
	query := "DELETE FROM user WHERE ID = 3 "
	res, err1 := db.Exec(query)
	if err1 != nil {
		panic(err1)
	}
	affectedRows, err2 := res.RowsAffected()
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("Del Success, The statement affected %d rows\n", affectedRows)
}
