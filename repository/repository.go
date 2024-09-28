package repository

import (
	"github.com/shimodii/go-crud/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
    Db *gorm.DB
}

var Database DbInstance

func OpenDatabase(){
    db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic("databased failed to open!")
    }

    db.AutoMigrate(&model.Card{})

    Database = DbInstance{
        Db: db,
    }

}
