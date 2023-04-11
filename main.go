// Package main is the entry point of the application
package main

import (
	"fmt"
	"net/http"
	"strings"
	"tennet/gethired/cmd"
	"tennet/gethired/infrastructure/repository/mysql"
	errorsController "tennet/gethired/infrastructure/restapi/controllers/errors"
	"tennet/gethired/infrastructure/restapi/middlewares"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"tennet/gethired/infrastructure/restapi/routes"
)

func main() {

	router := gin.Default()
	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())

	// mysql connection
	mysqlDB, err := mysql.NewGorm()
	if err != nil {
		_ = fmt.Errorf("fatal error in mysql file: %s", err)
		panic(err)
	}

	cmd.Execute(mysqlDB)

	router.Use(middlewares.GinBodyLogMiddleware)
	router.Use(errorsController.Handler)

	// mysql routes
	routes.ApplicationV1Router(router, mysqlDB)

	startServer(router)

}

func startServer(router http.Handler) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
		panic(err)

	}
	serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))
	s := &http.Server{
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		_ = fmt.Errorf("fatal error description: %s", strings.ToLower(err.Error()))
		panic(err)

	}
}
