package main

import (
	"strconv"

	"github.com/spf13/viper"
)

const defaultConfigPath = "./conf/conf.ini"

// AppConfigure server config.
type AppConfigure struct {
	Bind     string
	Port     int
	SavePath string
	Logger   struct {
		Path  string
		Level int
	}
	Proxy struct {
		Domain string
		Source string
	}
}

func mustAtoi(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v
}

func loadConfigure(path string) (*AppConfigure, error) {
	conf := &AppConfigure{}
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	conf.Bind = viper.Get("app.bind").(string)
	conf.Port = mustAtoi(viper.Get("app.port").(string))
	conf.SavePath = viper.Get("app.savepath").(string)
	conf.Logger.Path = viper.Get("logger.path").(string)
	conf.Logger.Level = mustAtoi(viper.Get("logger.level").(string))
	conf.Proxy.Domain = viper.Get("proxy.domain").(string)
	conf.Proxy.Source = viper.Get("proxy.source").(string)
	return conf, nil
}
