package main

import (
	"MelkOnline/internal/controller/ADregister"
	"MelkOnline/internal/controller/auth"
	"MelkOnline/internal/controller/chat"
	"MelkOnline/internal/controller/mainpage"
	"MelkOnline/internal/controller/signup"
	"os"

	model "MelkOnline/internal/core"

	"github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	isDBinitiated := os.Getenv("DB_INITIATED")
	if isDBinitiated == "false" {
		err = DB_init()
		if err != nil {
			panic(err)
		}
	}

	e := echo.New()

	e.Use(middleware.Logger())

	e.Static("/", "/home/ssaeidifarzad/ssfdata/ssaeidifarzad/Classes/S8/SE/Project/SE_project/public")
	e.File("/", "/home/ssaeidifarzad/ssfdata/ssaeidifarzad/Classes/S8/SE/Project/SE_project/public/index.html")

	e.POST("/signup", signup.NewSignupHandler().Signup)
	e.POST("/login", auth.NewLoginHandler().Login)
	e.POST("/ADregister", ADregister.NewADregisterHandler().ADregister)
	e.GET("/Chat/page/?chatid=", chat.NewChatHandler().GetMessagesByChatID)
	e.POST("/Chat/send/?chatid=", chat.NewChatHandler().SendMessage)
	e.GET("/mainpage", mainpage.NewMainpageHandler().GetAds) //??

	e.Logger.Fatal(e.Start(":8080"))
}

func DB_init() error {
	// This function is used to initialize the database connection
	// and create the tables if they do not exist.
	// It is called in the main function of the application.
	// The database connection is created using the GORM library.
	// The connection string is read from the environment variable DB_CONNECTION.
	// The tables are created using the AutoMigrate function of the GORM library.
	// The tables are created using the models
	// defined in the internal/core package.

	dbstr := os.Getenv("DB_CONNECTION")
	db, err := gorm.Open(
		mysql.Open(dbstr),
		&gorm.Config{})
	if err != nil {
		return err
	}
	DB, err := db.DB()
	if err != nil {
		return err
	}
	defer DB.Close()
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.AD{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Chat{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&model.Message{})
	if err != nil {
		return err
	}
	os.Setenv("DB_INITIATED", "true")
	return nil
}
