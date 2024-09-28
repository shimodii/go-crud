package service

import "github.com/shimodii/go-crud/model"

func ExistCheck(card model.Card) bool {
    if (card.ID == 0) {
        return true
    }
    return false
}
