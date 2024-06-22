package main

import (
	"log"

	"GoSnipeFun/bot"
	"GoSnipeFun/config"
)

func main() {
	switch config.StartingMode {
	case "bot":
		bot.StartBot(false)
	case "manager":
		bot.StartBot(true)
	case "cleaner":
		bot.StartCleaner()
	default:
		log.Fatalf("invalid starting mode: %s", config.StartingMode)
	}
}
