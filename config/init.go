package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"strings"
)

var (
	prefixEnvName = "__MOZZET"
	configPath    = ""
	file          = "mozzet.db"
)

func init() {
	initPath()

	_ = os.Setenv(EnvName("PATH"), configPath)

	fmt.Println(os.Getenv(EnvName("PATH")))
}

func EnvName(key string) string {
	return prefixEnvName + "_" + strings.ToUpper(key)
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
