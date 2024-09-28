package model

import "gorm.io/gorm"

type Card struct {
    gorm.Model
    Owner string
    Number string
}
