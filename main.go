package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type User struct {
	gorm.Model

	Nama     string `json:"nama" form:"nama"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	NoHP     string `json:"nohp" form:"nohP"`
	Saldo    uint32 `json:"saldo" form:"saldo"`
	Transfer []Transfer
	TopUp    []TopUp
}

type Transfer struct {
	gorm.Model

	NominalTransfer uint32 `json:"nominaltransfer" from:"nama"`
	NoHpTujuan      string `json:"nohptujuan" from:"nohptujuan"`
	NoHpPengirim    string `json:"nohppengirim" from:"nohppengirim"`
	UserID          uint
}

type TopUp struct {
	gorm.Model

	NominalTransfer uint32 `json:"nominaltransfer" from:"nama"`
	NoHpTujuan      string `json:"nohptujuan" from:"nohptujuan"`
	UserID          uint
}

func InitialMigration() {

	DB.AutoMigrate(Transfer{})
	DB.AutoMigrate(TopUp{})
	DB.AutoMigrate(User{})

}

func InitDB() {
	connectionString :=
		"root:admin123@tcp(localhost:3306)/service_app?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func Init() {
	InitDB()
	InitialMigration()
}

func main() {
	Init()
	fmt.Println("Account Service App")
	fmt.Println("=====================")
	fmt.Println("Menu")
	fmt.Println("No1. Registrasi")
	fmt.Println("No2. Tampilkan User")
	fmt.Println("No3. Update User")
	fmt.Println("No4. Delete User")
	fmt.Println("No5. Top Up")
	fmt.Println("No6. Transfer")
	fmt.Println("No7. History Top Up")
	fmt.Println("No8. History Transfer")
	fmt.Println("=====================")
	fmt.Println("Pilih Menu:")
	var menu string
	fmt.Scanln(&menu)

	switch menu {
	case "1":
	case "2":
	case "3":
	case "4":
	case "5":
	case "6":
	case "7":
	case "8":
	}

}
