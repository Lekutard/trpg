package db

import (
	"github.com/jinzhu/gorm"
	"github.com/trpg/infrasteucture/logging"
)

//DB DB Connections
type DB struct {
	MasterDB      *gorm.DB
	ReadReplicaDB *gorm.DB
}

//Database DBのインターフェース
type Database interface {
	Open() *DB
}

//Close DB Close
func (db *DB) Close() {
	db.MasterDB.Close()
	db.ReadReplicaDB.Close()
}

//LogMode Gorm LogMode
func (db *DB) LogMode(flg bool) {
	db.MasterDB.LogMode(flg)
	db.ReadReplicaDB.LogMode(flg)
}

//SetLogger Set Logger
func (db *DB) SetLogger(logger *logging.GormLogger) {
	db.MasterDB.SetLogger(logger)
	db.ReadReplicaDB.SetLogger(logger)
}
