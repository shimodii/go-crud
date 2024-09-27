package repository

import (
	"github.com/shimodii/go-crud/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase(){
    database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
    if err != nil {
        panic("databased failed to open!")
    }

    database.AutoMigrate(&model.Card{})
}
