package internal

import (
	"MelkOnline/internal/controller/ADregister"
	"MelkOnline/internal/controller/addnote"
	"MelkOnline/internal/controller/auth"
	"MelkOnline/internal/controller/chat"
	"MelkOnline/internal/controller/getpost"
	"MelkOnline/internal/controller/mainpage"
	"MelkOnline/internal/controller/searchfiltering"
	"MelkOnline/internal/controller/signup"
	"MelkOnline/internal/controller/signup/payment"
	"MelkOnline/internal/controller/user"
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
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowMethods, echo.HeaderAccessControlAllowCredentials},
	}))
	es.e.POST("api/v1/signup/", signup.NewSignupHandler().Signup)
	es.e.GET("api/v1/signup/payment/", payment.NewPaymentHandler().Pay)
	es.e.POST("api/v1/login/", auth.NewLoginHandler().Login)
	es.e.POST("api/v1/ads/register", ADregister.NewADregisterHandler().ADregister)
	es.e.GET("api/v1/ads/:ad_id/chats", chat.NewChatHandler().GetMessage)
	es.e.POST("api/v1/ads/:ad_id/chats", chat.NewChatHandler().SendMessage)
	es.e.GET("api/v1/ads/mainpage/", mainpage.NewMainpageHandler().GetAds)
	es.e.GET("api/v1/ads/searchfiltering/", searchfiltering.NewSearchFilteringHandler().Searchfiltering)
	es.e.GET("api/v1/ads/:ad_id", getpost.NewGetPostHandler().GetPost)
	es.e.GET("/swagger/*", echoSwagger.WrapHandler)
	es.e.GET("api/v1/ads/:ad_id/notes", addnote.NewNoteHandler().GetNotes)
	es.e.POST("api/v1/ads/:ad_id/notes", addnote.NewNoteHandler().StoreNote)
	es.e.GET("api/v1/user", user.NewUserHandler().GetUserBySession)
	es.e.GET("api/v1/user/ads", user.NewUserHandler().GetUserAds)
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
	err = db.AutoMigrate(&model.Payment{})
	if err != nil {
		return err
	}
	err = db.Exec("CREATE TABLE IF NOT EXISTS images (ID INT AUTO_INCREMENT PRIMARY KEY, AD_ID INT, path VARCHAR(255))").Error
	if err != nil {
		return err
	}
	os.Setenv("DB_INIT", "true")
	return nil
}
