package env

import (
	"os"
	"strings"
)

type Env struct {
	name string
}

func (e Env) GetEnv(key string) string {
	return os.Getenv(e.envName(key))
}

func (e Env) SetEnv(key string, value string) {
	_ = os.Setenv(e.envName(key), value)
}

func (e Env) envName(key string) string {
	return e.name + "_" + strings.ToUpper(key)
}

func New(name string) Env {
	return Env{name: name}
}
