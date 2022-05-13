package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"tg-monitor/internal/telegram"
	webserver "tg-monitor/internal/webServer"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
}
func main() {
	go telegram.StartTelegram()
	go webserver.Listen()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shuting down Server ...")
}
