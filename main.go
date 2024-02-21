package main

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Table1 struct {
	UserName string `gorm:"primaryKey"`
	UserUUID string
	Ak       string
	KmsName  string
}

type Table2 struct {
	KeyName string `gorm:"primaryKey"`
	Access  [][]byte
}

func init() {
	// prepare
}

func main() {
	sqliteDB, err := gorm.Open(sqlite.Open("/tmp/test.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	err = sqliteDB.AutoMigrate(
		// Table-1
		&Table1{},
		// Table-2
		&Table2{},
	)

	if err != nil {
		log.Panic(err)
	}

	table1 := Table1{
		UserName: "test_register_info",
		UserUUID: uuid.NewString(),
		Ak:       "abfeeee",
		KmsName:  "vault",
	}

	ret := sqliteDB.Create(&table1)
	if ret.Error != nil {
		log.Panic(ret.Error)
	}

	var regInfo Table1
	ret = sqliteDB.Find(&regInfo, Table1{UserUUID: table1.UserUUID})
	if ret.Error != nil {
		log.Panic(ret.Error)
	}

	if regInfo.Ak == table1.Ak {
		log.Println("query ok")
	} else {
		log.Println("query failed")
	}

	updates := map[string]interface{}{
		"KmsName": "aws-kms",
		"Ak":      "fffeee",
	}
	sqliteDB.Model(&regInfo).Select("*").Updates(updates)

	table2 := Table2{
		KeyName: "key-name-1",
		Access:  [][]byte{[]byte("this is access")},
	}

	ret2 := sqliteDB.Create(&table2)
	if ret2.Error != nil {
		log.Panic(ret.Error)
	}

}
