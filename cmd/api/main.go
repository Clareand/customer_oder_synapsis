package main

import (
	"fmt"

	"github.com/Clareand/customer_oder_synapsis/config/postgresql"

	"net/http"
	"os"

	authHandler "github.com/Clareand/customer_oder_synapsis/pkg/auth/handler"
	authRepo "github.com/Clareand/customer_oder_synapsis/pkg/auth/repository"
	authUsecase "github.com/Clareand/customer_oder_synapsis/pkg/auth/usecase"

	configRedis "github.com/Clareand/customer_oder_synapsis/config/redis"

	echoPrometheus "github.com/globocom/echo-prometheus"
	config "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmechov4/v2"
	"go.elastic.co/apm/module/apmlogrus/v2"
)

var fileName = make(chan string)

func main() {
	if os.Getenv("GO_ENV") == "local" {
		if err := config.Load(".env"); err != nil {
			fmt.Println(".env is not loaded properly")
			fmt.Println(err)
			os.Exit(2)
		}
	}

	dbConn := postgresql.CreateConnection()
	redisConn := configRedis.CreateConnection()

	logrus.AddHook(&apmlogrus.Hook{
		LogLevels: []logrus.Level{
			logrus.ErrorLevel,
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.InfoLevel,
		},
	})

	r := echo.New()
	r.Debug = true
	r.Use(echoPrometheus.MetricsMiddleware())
	r.Use(apmechov4.Middleware())
	r.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	r.Use(middleware.Recover())
	r.Use(middleware.Logger())
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Type", "Authorization", "traceparent", "tracestate", "va-key"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	r.GET("/", func(context echo.Context) error {
		return context.HTML(http.StatusOK, "<strong>System Configuration "+os.Getenv("APP_VERSION")+"</strong>")
	})

	apiV1 := r.Group("/api/v1")

	authRepo := authRepo.NewLoginRepo(dbConn, redisConn)
	authUsecase := authUsecase.NewLoginRepo(authRepo)
	authHandler.NewHTTPHandler(authUsecase).Mount(apiV1)

	err := r.Start(":" + os.Getenv("PORT"))
	if err != nil {
		log.Error().Msg(err.Error())
	}
}
