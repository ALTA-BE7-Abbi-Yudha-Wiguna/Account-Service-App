package controllers

import (
	"fmt"
	"gorm.io/gorm"
	"withgorm/entities"
)

func HistoriTransfer(db *gorm.DB) ([]entities.Transfer, error) {
	var noHp string
	fmt.Println("Masukan No.Hp:")
	fmt.Scanln(&noHp)

	var transfer []entities.Transfer
	err := db.Where("no_hp_pengirim = ?", noHp).Find(&transfer).Error
	if err != nil {
		// panic(tx.Error)
		return transfer, err
	}
	return transfer, nil
}
