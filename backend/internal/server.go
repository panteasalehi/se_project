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
	"MelkOnline/internal/infrastructure"
	"net/http"
	"os"

	_ "MelkOnline/docs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type EchoServer struct {
	e *echo.Echo
}

func NewEchoServer() *EchoServer {
	echo := echo.New()
	es := &EchoServer{e: echo}
	err := godotenv.Load("/home/ssaeidifarzad/ssfdata/ssaeidifarzad/Classes/S8/SE/Project/SE_project/backend/.env")
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
		AllowOrigins:     []string{"http://45.147.97.39:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowMethods, echo.HeaderAccessControlAllowCredentials},
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

func (es *EchoServer) Start(port string) {
	isDBinitiated := os.Getenv("DB_INIT")
	if isDBinitiated == "" {
		panic("DB_INIT is not set")
	} else if isDBinitiated == "false" {
		err := DB_init()
		if err != nil {
			panic(err)
		}
	}
	es.e.Logger.Fatal(es.e.Start(":" + port))
}

func (es *EchoServer) Stop() {
	es.e.Close()
}

func DB_init() error {
	db := infrastructure.GetDB()
	err := db.AutoMigrate(&model.User{})
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
