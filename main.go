package main

import (
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/database"
	"github.com/nandaditra/task-5-vix-btpns-Nanda-Aditya-Putra/router"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.DatabaseConnection()
)

func main() {
	defer database.CloseDatabaseConnection(db)

	router.AuthRouter()
	router.PhotoRouter()
	router.UserRouter()
}
