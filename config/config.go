package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
)

var (
	Path = ""
	file = "config.json"
)

func init() {
	home, err := homedir.Dir()

	if err != nil {
		fmt.Println("Home Path Error: " + err.Error())
		os.Exit(1)
	}

	Path = home + "/.mozzet"
}

func InitConfig() {
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		initConfigPath()
	}

	if _, err := os.Stat(Path + "/" + file); os.IsNotExist(err) {
		initConfigFile()
	}

	viper.AddConfigPath(Path)
	viper.SetConfigName(file)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func initConfigPath() {
	err := os.MkdirAll(Path, os.ModePerm)

	if err != nil {
		fmt.Println("Config Path Create Err:" + err.Error())
		os.Exit(1)
	}
}

func initConfigFile() {

}
