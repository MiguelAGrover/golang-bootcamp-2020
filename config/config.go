package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

// config : structure that contain port to be used by server, pages from which data is retrieved and csv paths.
type config struct {
	Server struct {
		Address string
	}
	Sources struct {
		DigimonAPI string
	}
	Dest struct {
		DigimonCSV string
	}
}

// C : instance of the config structure
var C config

// ReadConfig : loads configuration from config.yml and loads it into a config structure.
func ReadConfig() {
	Config := &C

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join("$GOPATH", "src",
		"github.com", "MiguelAGrover",
		"golang-bootcamp-2020", "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	spew.Dump(C)
}
