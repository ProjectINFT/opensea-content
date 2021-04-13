package main

import "testing"

func TestConfig(t *testing.T) {
	if conf, err := loadConfigure(defaultConfigPath); err != nil {
		t.Error("loadConfigure err, got ", err)
	} else {
		if conf.Bind != "0.0.0.0" || conf.Port != 8008 || conf.Logger.Path != "./logs" || conf.Logger.Level != 3 || conf.Proxy.Domain != "baidu.com" || conf.Proxy.Source != "google.com" {
			t.Error("loadConfigure value ")
		}
	}
}
