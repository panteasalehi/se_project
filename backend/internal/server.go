package internal

import (
	"MelkOnline/internal/controller/ADregister"
	"MelkOnline/internal/controller/auth"
	"MelkOnline/internal/controller/chat"
	"MelkOnline/internal/controller/getpost"
	"MelkOnline/internal/controller/mainpage"
	"MelkOnline/internal/controller/searchfiltering"
	"MelkOnline/internal/controller/signup"
	model "MelkOnline/internal/core"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type EchoServer struct {
	e *echo.Echo
}

func NewEchoServer() *EchoServer {
	echo := echo.New()
	es := &EchoServer{e: echo}
	err := godotenv.Load("/home/runner/work/se_project/se_project/backend/.env")
	if err != nil {
		panic(err)
	}
	es.Route()
	return es
}

func (es *EchoServer) Route() {
	es.e.Use(middleware.Logger())
	es.e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	es.e.POST("/signup", signup.NewSignupHandler().Signup)
	es.e.POST("/login", auth.NewLoginHandler().Login)
	es.e.POST("/ADregister", ADregister.NewADregisterHandler().ADregister)
	es.e.GET("/Chat/page/?chatid=", chat.NewChatHandler().GetMessagesByChatID)
	es.e.POST("/Chat/send/?chatid=", chat.NewChatHandler().SendMessage)
	es.e.GET("/mainpage", mainpage.NewMainpageHandler().GetAds)
	es.e.GET("/searchfiltering", searchfiltering.NewSearchFilteringHandler().Searchfiltering)
	es.e.GET("/getpost/:id", getpost.NewGetPostHandler().GetPost)
	es.e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func (es *EchoServer) Start() {
	isDBinitiated := os.Getenv("DB_INIT")
	fmt.Println(isDBinitiated)
	if isDBinitiated == "" {
		panic("DB_INIT is not set")
	} else if isDBinitiated == "false" {
		err := DB_init()
		if err != nil {
			panic(err)
		}
	}
	es.e.Logger.Fatal(es.e.Start(":8080"))
}

func (es *EchoServer) Stop() {
	es.e.Close()
}

func DB_init() error {
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
	os.Setenv("DB_INIT", "true")
	return nil
}
