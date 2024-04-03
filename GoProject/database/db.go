package database

import (
	"fmt"

	m "GoProject/eventhandler"

	mod "GoProject/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB() {
	dsn := "root:Sathyabama*40110529@tcp(localhost:3306)/report?parseTime=true"
	// dsn := "user:pass@tcp(db:3306)/product?parseTime=true"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("Unable to Connect")
	}
	//Table creation for Namespaces.
	errr := db.AutoMigrate(&mod.Namespace{})
	if errr != nil {
		fmt.Println("Error migrating database:", errr)
		panic("Unable to migrate database")
	}

	m.SetDB(db)
}
