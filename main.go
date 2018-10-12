package main

import (
	"sms/server"
)

func main() {
	opts := &server.Options{}

	opts, err := server.ProcessConfigFile(server.DEFAULT_CONFIG_FILE)
	if err != nil {
		server.PrintAndDie("load config fail:" + err.Error())
	}
	s := server.New(opts)
	s.ConfigureLogger()
	s.Start()
}
