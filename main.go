package main

import (
	"github.com/vi350/vk-internship/internal/app"
	"log"
)

func main() {
	bot, err := app.New()
	if err != nil {
		log.Panic(err)
	}

	err = bot.Run()
	if err != nil {
		log.Panic(err)
	}
}
