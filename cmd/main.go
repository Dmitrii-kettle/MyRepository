package main

import (
	"MyServer"
	"log"
)

func main() {
	srv := new(poj.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Произошла ошибка при запуске HTTP сервера: %s", err.Error())
	}
}
