package db

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Mysql dbをmysqlにする場合はこれを使う
type Mysql struct {
}

//NewMysql New Mysql
func NewMysql() *Mysql {
	return &Mysql{}
}

//Open MySQL Open
func (mysql *Mysql) Open() *DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	option := "charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo&timeout=5s" //5秒でタイムアウト

	//master
	masterPort := os.Getenv("DB_MASTER_PORT")
	masterDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, masterPort, name, option)
	masterDB, err := gorm.Open("mysql", masterDsn)
	if err != nil {
		panic(err)
	}
	masterDB.DB().SetConnMaxLifetime(time.Hour)
	masterDB.DB().SetMaxIdleConns(5)
	masterDB.DB().SetMaxOpenConns(100)

	//readReplica
	readReplicaPort := os.Getenv("DB_READ_REPLICA_PORT")
	readReplicaDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, readReplicaPort, name, option)
	readReplicaDB, err := gorm.Open("mysql", readReplicaDsn)
	if err != nil {
		panic(err)
	}
	readReplicaDB.DB().SetConnMaxLifetime(time.Hour)
	readReplicaDB.DB().SetMaxIdleConns(10)
	readReplicaDB.DB().SetMaxOpenConns(500)

	return &DB{MasterDB: masterDB, ReadReplicaDB: readReplicaDB}
}
