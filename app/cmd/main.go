package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	_handler "github.com/gamepkw/test-argo/app/internal/handler"
)

func init() {
	viper.SetConfigFile(`../../config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	e := echo.New()
	_handler.NewHandler(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
