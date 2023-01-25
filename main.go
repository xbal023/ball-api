package main

import (
	"os"
	"sync"
	"time"
	"os/exec"
	"syscall"
	"os/signal"
	
	"github.com/joho/godotenv"
	"github.com/briandowns/spinner"

	"github.com/bolaxd/dumn/app"
	"github.com/bolaxd/dumn/config"
	wa "github.com/bolaxd/dumn/bot.wa"
	)

func main()  {
	var wg sync.WaitGroup
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	s := spinner.New(config.SpinnerSet, 121*time.Millisecond)
	s.Prefix = "Progress Running App...\n\n"
	s.Color("red")
	s.Start()
	time.Sleep((37 * 126) *time.Millisecond)
	s.Stop()
	wg.Add(2)
	go wa.RunBot(&wg)
	go app.Run(&wg)
	// fmt.Println("2 Aplikasi sedang dijalankan secara bersamaan, wait for Running...")
	wg.Wait()
	/**Aplikasi crash Tidak akan tertutup paksa**/
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}