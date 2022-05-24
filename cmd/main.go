package main

import (
	poj "MyServer"
	"MyServer/config"
	"log"
)

func main() {
	srv := new(poj.Server)
	if err := srv.Run(config.GetConfigs()); err != nil {
		log.Fatalf("Произошла ошибка при запуске HTTP сервера: %s", err.Error())
	}
}
