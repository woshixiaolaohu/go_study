package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string `gorm:"<-:create"`
	Age          uint8
	Email        *string
	Birthday     *time.Time
	MemberNumber sql.NullString
}
type Blog struct {
	ID      int
	Author  User `gorm:"embedded"`
	Upvotes int32
}

func init() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err1 := sql.Open("mysql", "db_test_1")
	if err1 != nil {
		panic(err1)
	}
	gormDB, err2 := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("dormDB type is %T,%T\n", gormDB, db)
}
