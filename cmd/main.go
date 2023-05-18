package main

import (
	"log"

	"github.com/cocoide/transaction/pkg/domain/model"
	"github.com/cocoide/transaction/pkg/domain/repo"
	"github.com/cocoide/transaction/pkg/interface/handler"
	"github.com/cocoide/transaction/pkg/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	DSN := "kazuki:secret@tcp(db:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(DSN))
	if err != nil {
		log.Fatalf("failed to connect with databse: %s", err.Error())
	}
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("failed to migrate database: %s", err.Error())
	}
	tx := repository.NewTxRepo(db)
	ur := repository.NewUserRepo(db)
	au := usecase.NewAuthUseCase(ur, tx)
	uh := handler.NewAuthHandler(au)

	e.POST("/user", uh.Register)
	e.Logger.Fatal(e.Start(":8080"))
}
