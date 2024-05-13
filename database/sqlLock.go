package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

/*
*
分布式读写锁

	原则
		1：读模式共享，写模式互斥
		2：存在三种模式状态：读加锁状态、写加锁状态、无锁状态
*/

var (
	db         *gorm.DB
	dbUserName = "root"
	dbPassWord = "123456"
	dbHost     = "127.0.0.1:3306"
	dbDatabase = "db_test_1"

	stateReadLock  = "ReadLock"
	stateWriteLock = "WriteLock"
	stateUnLock    = "Unlock"
)

type RWLock struct {
	LockMark      string `gorm:"default:'Unlock'"`
	ReadLockCount uint32 `gorm:"default:0"`
	LockReason    string
}

type Stock struct {
	gorm.Model
	RWLock
	Count int64
}

func (Stock) TableName() string {
	return "stocks"
}

// 读锁模式
func (s Stock) Rlock(db *gorm.DB, lockReason string) error {
	condition := "(id = ?) AND (lock_mark != ?)"
	fields := map[string]interface{}{
		"lock_mark":       stateReadLock,
		"read_lock_count": gorm.Expr("read_lock_count + ?", 1),
		"lock_reason":     lockReason,
	}
	result := db.Model(&Stock{}).Where(condition, s.ID, stateWriteLock).Updates(fields)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to rlock Stock, RowsAffected=0")
	}
	return nil
}

func (s Stock) RUnlock(db *gorm.DB, UnLockReason string) error {
	query := fmt.Sprintf(`UPDATE stocks SET read_lock_count=if(read_lock_count>0,read_lock_count-1,0), lock_mark=if(read_lock_count<1, 'Unlock', 'ReadLock'),lock_reason ='%s' where id= %d and lock_mark='%s'`, UnLockReason, s.ID, stateReadLock)
	result := db.Exec(query)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to RUnlock Stock, RowsAffected=0")
	}
	return nil
}

func init() {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUserName, dbPassWord, dbHost, dbDatabase)
	//mysqlConfig := mysql.Config{DSN: dsn}
	//gormConfig := &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}
	//
	//var err error
	//if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
	//	panic(err)
	//}
	//db.Set("db:table_options", "ENGINE = InnoDB DEFAULT CHARSET = utf8")
	//// register tables
	//if err = db.AutoMigrate(&Stock{}); err != nil {
	//	panic(err)
	//}
	//if result := db.Model(&Stock{}).Save(&Stock{Model: gorm.Model{}, RWLock: RWLock{}, Count: 10}); result.Error != nil {
	//	panic(result.Error)
	//}
	//fmt.Printf("db type %T", db)
}
