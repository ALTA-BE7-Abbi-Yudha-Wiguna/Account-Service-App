package controllers

import (
	"fmt"
	"gorm.io/gorm"
	"withgorm/entities"
)

func HistoriTopup(db *gorm.DB) ([]entities.TopUp, error) {
	var noHp string
	fmt.Println("Masukan No.Hp:")
	fmt.Scanln(&noHp)

	var topups []entities.TopUp
	err := db.Where("no_hp_tujuan = ?", noHp).Find(&topups).Error
	if err != nil {
		// panic(tx.Error)
		return topups, err
	}
	return topups, nil
}
