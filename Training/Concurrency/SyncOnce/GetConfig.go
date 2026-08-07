package main

import "sync"

var (
	mx     = sync.Mutex{}
	config *Config
)

func GetConfig() *Config {

	if config == nil {

		mx.Lock()
		defer mx.Unlock()

		config = &Config{}
	}
	return config

}
