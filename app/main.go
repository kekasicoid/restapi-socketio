package main

import (
	"strconv"
	"strings"
	"time"

	_keluargaHandler "github.com/kekasicoid/restapi-socketio/controller/api/keluarga/handler"
	_keluargaRepo "github.com/kekasicoid/restapi-socketio/controller/api/keluarga/repository"
	_keluargaUC "github.com/kekasicoid/restapi-socketio/controller/api/keluarga/usecase"
	_keluargaSioHandler "github.com/kekasicoid/restapi-socketio/controller/socketio/keluarga/handler"
	"github.com/kekasicoid/restapi-socketio/helper"
	docs "github.com/kekasicoid/restapi-socketio/swagger"
	"github.com/sirupsen/logrus"
	"go.uber.org/ratelimit"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/kekasicoid/kekasigohelper"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
)

func init() {
	viper.SetConfigFile(`./.env`)
	err := viper.ReadInConfig()
	if err != nil {
		kekasigohelper.LoggerFatal(err)
	}
}

// @title RestAPI & Socket.io v2
// @version 1.0
// @description Pattern Go RestAPI + Socket.io v2. with the same ports.
// @termsOfService https://id.linkedin.com/public-profile/in/arditya-kekasi

// @contact.name Arditya Kekasi
// @contact.url http://www.kekasi.co.id
// @contact.email arditya@kekasi.co.id

// @license.name YouTube KekasiGen
// @license.url https://kekasi.link/kekasigensub
// @schemes http https
var (
	limit ratelimit.Limiter
	count = 0
	//rps   = flag.Int("rps", 100, "request per second")
)

func leakBucket() gin.HandlerFunc {
	count++
	prev := time.Now()
	return func(c *gin.Context) {
		now := limit.Take()
		kekasigohelper.LoggerWarning(now.Sub(prev))
		prev = now
	}
}

func main() {
	docs.SwaggerInfo.BasePath = ""
	redisPort := viper.Get("REDIS_PORT").(string)
	redisAddress := viper.Get("REDIS_ADDRESS").(string)
	redisPassword := viper.Get("REDIS_PASSWORD").(string)
	port := viper.Get("APP_PORT").(string)
	appmode := viper.Get("APP_MODE").(string)
	appTimeout := viper.GetInt("APP_TIMEOUT")
	timeoutContext := time.Duration(appTimeout) * time.Second
	validate := helper.ValidatorInit()

	if viper.Get("GRAYLOG_ACTIVE").(string) == "on" {
		// Create a new Graylog hook
		hook := graylog.NewGraylogHook(viper.Get("GRAYLOG_URL").(string), nil)
		logrus.AddHook(hook)
		logrus.Info("[HTTP] RestAPI-Socketio IS RUNNING")
	}

	db := helper.DBMysqlConn()
	helper.DbMirgrator(db)

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	server := socketio.NewServer(nil)

	_, err := server.Adapter(&socketio.RedisAdapterOptions{
		Host:     redisAddress,
		Addr:     redisAddress + ":" + redisPort,
		Port:     redisPort,
		Prefix:   "kekasigen",
		Password: redisPassword,
	})

	if err != nil {
		kekasigohelper.LoggerFatal("Error server.Adapter :" + err.Error())
	}
	middL := helper.InitMiddleware()
	r.Use(gin.Recovery())
	//r.Use(leakBucket())

	origin := strings.Split(viper.Get("ALLOW_ORIGIN").(string), ",")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     origin,
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.SetTrustedProxies(nil)
	if appmode == "development" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	kks := 0
	r.GET("/ping", middL.LeakBucketUber(), func(c *gin.Context) {
		kks++
		c.JSON(200, gin.H{
			"message": "Hello World Ping" + strconv.Itoa(kks),
		})
		return
	})
	kks2 := 0
	r.GET("/pong", middL.LimitRequest(), func(c *gin.Context) {
		kks2++
		c.JSON(200, gin.H{
			"message": "Hello World Pong" + strconv.Itoa(kks2),
		})
		return
	})

	go func() {
		if err := server.Serve(); err != nil {
			kekasigohelper.LoggerFatal("socketio listen error: " + err.Error())
		}
	}()
	defer server.Close()

	// init repositories
	keluargaRepo := _keluargaRepo.NewKeluargaRepository(db)

	// init Usecase
	keluargaUC := _keluargaUC.NewKeluargaUsecase(keluargaRepo, timeoutContext)

	// ini Handler
	_keluargaHandler.NewKeluargaHandler(r, keluargaUC, validate, server)
	_keluargaSioHandler.NewKeluargaSioHandler(r, keluargaUC, nil, validate, server, middL)

	err = r.Run(":" + port)
	if err != nil {
		kekasigohelper.LoggerFatal(err)
	}
}
