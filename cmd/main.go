package main

import (
	"os"
	"sync"

	"multiplica/cmd/config"
	"multiplica/internal/handler/http"
	"multiplica/internal/handler/grpc"

	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetFormatter(&log.TextFormatter{
        FullTimestamp: true,
    })

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("cmd/config/")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration config.Configurations

	if err := viper.ReadInConfig(); err != nil {
		log.WithFields(log.Fields{
            "Error": err.Error(),
        }).Fatal("Failed to listen")

	  os.Exit(2)
	}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.WithFields(log.Fields{
            "Error": err.Error(),
        }).Fatal("Failed to listen")

		os.Exit(2)
	}

	var wg sync.WaitGroup

	wg.Add(1)
  //start Rest api Server
  go http.NewHTTPHandler(configuration)

	wg.Add(1)
	go grpc.NewGRPCHandler(configuration)

	wg.Wait()
}
