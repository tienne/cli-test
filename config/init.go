package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mozzet/cli/database"
	"github.com/mozzet/cli/env"
	"os"
)

var (
	enviroment env.Env
	configPath = ""
	file       = "mozzet.db"
)

func init() {
	initPath()

	enviroment = env.New("MOZZET")
	enviroment.SetEnv("home", configPath)
}

func initPath() {
	home, err := homedir.Dir()

	if err != nil {
		fmt.Println("Home configPath Error: " + err.Error())
		os.Exit(1)
	}

	configPath = home + "/.mozzet"
}

func InitConfig() {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		initConfigPath()
	}

	if _, err := os.Stat(configPath + "/" + file); os.IsNotExist(err) {
		initDatabase()
	}
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

func initDatabase() {
	err := database.Execute(func(db *database.DB) error {
		return db.Set("setting", "current.version", "aaa")
	})

	fmt.Println(err)
}
