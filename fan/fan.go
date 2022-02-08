package fan

import (
	"fmt"
	"gitee.com/hilaoyu/go-basic-utils/utilCmd"
	"gitee.com/hilaoyu/go-basic-utils/utilFile"
	"github.com/stianeikeland/go-rpio/v4"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
	"time"
)

var cpuTemperature float32 = 0
var cpuTemperatureLock sync.Mutex

func CpuTemperatureWrite(t float32) {
	cpuTemperatureLock.Lock()
	cpuTemperature = t
	cpuTemperatureLock.Unlock()
}

func CpuTemperatureRead() float32 {
	cpuTemperatureLock.Lock()
	t := cpuTemperature
	cpuTemperatureLock.Unlock()
	return t
}

func ServRun() {
	i := 1
	for {
		FanRun()
		fmt.Println("fan.FanRun exit: ", i)
		time.Sleep(time.Duration(5) * time.Second)
		i++
	}
}
func FanRun() {
	if !utilFile.Exists("/dev/mem") && utilFile.Exists("/dev/gpiomem") {
		utilCmd.RunCommand(true, "ln", "-s", "/dev/gpiomem", "/dev/mem")

	}
	err := rpio.Open()
	if nil != err {
		fmt.Println("rpio.Open err: ", err)
	}
	defer rpio.Close()
	pin := rpio.Pin(4)
	pin.Output()

	//pin.Toggle()

	for {
		tmpCore, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
		if nil != err {
			fmt.Println("err read:", err)
			tmpCore = []byte("48000")
		}

		tmpCoreInt, err := strconv.Atoi(strings.TrimSpace(string(tmpCore)))
		if nil != err {
			fmt.Println("err to int:", err)
		}

		CpuTemperatureWrite(float32(tmpCoreInt/10) / 100)

		if tmpCoreInt > 47000 {
			pin.High()
		} else if tmpCoreInt < 44000 {
			pin.Low()
		}

		time.Sleep(time.Duration(1) * time.Second)
	}

}
