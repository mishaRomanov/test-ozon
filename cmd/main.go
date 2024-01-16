package main

import (
	"database/sql"
	"fmt"
	"github.com/caarlos0/env/v10"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mishaRomanov/test-ozon/config"
	"github.com/mishaRomanov/test-ozon/internal/handler"
	"github.com/mishaRomanov/test-ozon/internal/storage"
	cache "github.com/mishaRomanov/test-ozon/internal/storage/cache"
	"github.com/mishaRomanov/test-ozon/internal/storage/postgres"
	"github.com/sirupsen/logrus"
)

func InitConfig() (config.Config, error) {
	var cfg config.Config
	//here we parse all environmental variables into struct object
	err := env.Parse(&cfg)
	if err != nil {
		return config.Config{}, err
	}
	return cfg, nil
}

// func that creates storage depending on an environmental variable
func createStorageBasedOnFlag(config string, db *sql.DB) storage.Storager {
	if config == "postgres" || config == "Postgres" {
		logrus.Infoln("Creating Postgres database...")
		return postgres.Create(db)
	}
	if config == "cache" || config == "Cache" {
		logrus.Infoln("Creating in-memory storage...")
		db.Close()
		return cache.NewCache()
	}
	logrus.Infoln("No correct config option detected. Creating in-memory storage...")
	return cache.NewCache()
}

func main() {
	logrus.Infoln(
		`[DEBUG] Handlers available: 
				POST --> /link/add
				GET --> /link/:shortened_link
				GET --> /about`)
	gin.SetMode(gin.ReleaseMode)
	//create a server
	service := gin.Default()

	//if enabled, client IP will be parsed from the request's headers that
	// match those stored at `(*gin.Engine).RemoteIPHeaders`. If no IP was
	// fetched, it falls back to the IP obtained from
	// `(*gin.Context).Request.RemoteAddr`.
	service.ForwardedByClientIP = true
	err := service.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		logrus.Fatalf("%v/n", err)
	}
	//setting up config
	cfg, err := InitConfig()
	if err != nil {
		logrus.Fatal("error while setting up config. Check environmental variables and try again")
	}

	//creating a connection string which has all the needed information:
	//user, db name, address,password
	connectString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", cfg.User, cfg.Password, cfg.Address, cfg.Port, cfg.DatabaseName)
	//creating database object
	database, err := sql.Open("postgres", connectString)
	if err != nil {
		logrus.Fatalf("Failed to open database: %v", err)
	}

	//creating a handler.Handler object which contains needed storage type
	//the func createStorageBasedOnFlag returns a storage type specified by environmental variable
	handlerObject := handler.New(createStorageBasedOnFlag(cfg.Storage, database))

	//endpoint returns the full link if found
	//the short one is given through :shortLink parameter
	service.GET("/link/:shortLink", handlerObject.HandleGet)

	//endpoint creates and returns a new shortened link
	//the original one is sent through json
	service.POST("/link/add", handlerObject.HandlePost)

	//endpoint that handles /about requests
	service.GET("/about", func(ctx *gin.Context) {
		str := `
Hello! 
This is a small service that creates short links of your old long links.
Send a POST request to /link/add with your link in request body like json {"url":"your_link_here"} to receive a new one
And then send a GET request to /link/*put_your_link_here* to receive your full link.
`
		ctx.String(http.StatusOK, str)
	})

	//start listening
	logrus.Fatalf("%v", service.Run(":8080"))

}
