package config

import (
    "github.com/spf13/viper"
    "fmt"
)

var conf *viper.Viper

func init() {
    conf = viper.New()
    conf.SetConfigName("config")
    conf.SetConfigType("json")
    conf.AddConfigPath("/etc/honeystart/")
    conf.AddConfigPath(".")
    err := conf.ReadInConfig()
    if err != nil {
        // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
    }
}

func Secret() string {
    return conf.GetString("server.secret")
}

func SecretByte() []byte {
    return []byte(Secret())
}