package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path"
)

var (
	configPath = ""
	file       = "config"
	ext        = "yaml"
)

func initPath() {
	home, err := homedir.Dir()

	if err != nil {
		fmt.Println("Home configPath Error: " + err.Error())
		os.Exit(1)
	}

	configPath = home + "/.mozzet"
}

func InitConfig() {
	initPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		initConfigPath()
	}

	if _, err := os.Stat(configPath + "/" + file); os.IsNotExist(err) {
		initConfigFile()
	}
	viper.SetConfigType(ext)
	viper.AddConfigPath(configPath)
	viper.SetConfigName(file)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	viper.SetDefault("aaa", 123)

	var config Config
	viper.Unmarshal(&config)
	fmt.Printf("%#v", config)
}

func ConfigPath() string {
	return configPath
}

func initConfigPath() {
	err := os.MkdirAll(configPath, os.ModePerm)

	if err != nil {
		fmt.Println("Config configPath Create Err:" + err.Error())
		os.Exit(1)
	}
}

func initConfigFile() {
	err := ioutil.WriteFile(configFilePath(), []byte(""), 0644)

	if err != nil {
		fmt.Println("Config File Not Create Err:" + err.Error())
		os.Exit(1)
	}
}

func configFilePath() string {
	return path.Join(configPath, file+"."+ext)
}
