package model

type Card struct {
    ID uint `json:"id" gorm:"primaryKey"`
    Owner string `json:"owner"`
    Number string `json:"number"`
}
