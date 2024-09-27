package repository

import (
	"github.com/shimodii/go-crud/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func OpenDatabase(){
    Database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic("databased failed to open!")
    }

    Database.AutoMigrate(&model.Card{})

}
