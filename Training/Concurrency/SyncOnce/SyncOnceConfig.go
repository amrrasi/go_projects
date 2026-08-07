package main

import "sync"

var once sync.Once

func GetConfOnce() *Config {

	if config == nil {

		once.Do(
			func() {

				println("Making a Config ...")
				config = &Config{}

			},
		)
	}
	println("Using Old Conf")
	return config

}
