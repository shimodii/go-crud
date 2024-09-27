package model

import "gorm.io/gorm"

type Card struct {
    gorm.Model
    Id string
    Owner string
    Number string
}
