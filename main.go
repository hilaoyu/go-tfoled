package main

import (
	"go_tfoled/fan"
	"go_tfoled/led"
)

func main() {

	go fan.ServRun()
	led.ServRun()

	/*for {
		time.Sleep(time.Duration(10) * time.Second)
	}*/
}
