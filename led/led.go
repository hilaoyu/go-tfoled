package led

import (
	"fmt"
	"gitee.com/hilaoyu/go-basic-utils/utils"
	"github.com/mdp/smallfont"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"go_tfoled/fan"
	"go_tfoled/i2c"
	"gobot.io/x/gobot/platforms/raspi"
	"image"
	"image/color"
	"time"
)

func ServRun() {
	for {
		Show()
		time.Sleep(time.Duration(5) * time.Second)
	}
}
func Show() {
	width := 128
	height := 32

	//utilImg.SaveImage(filepath.Join(utils.GetSelfPath(), "aa.png"), img)
	//return
	r := raspi.NewAdaptor()
	oled := i2c.NewSSD1306Driver(r, i2c.WithSSD1306DisplayWidth(width), i2c.WithSSD1306DisplayHeight(height))

	defer func() {
		oled.Off()
		Show()
	}()

	err := oled.Start()
	if err != nil {
		fmt.Println("oled start:", err)
		return
	}
	viewCacheChan := make(chan string, 10)
	work := func() {

		oled.Clear()
		img := image.NewRGBA(image.Rect(0, 0, width, height))

		i := 0
		cpuUsage, err := cpu.Percent(time.Second, false)
		if err == nil {
			var cpuAllUse float64 = 0
			for _, cu := range cpuUsage {
				cpuAllUse += cu
			}

			textToImage("CPU:"+fmt.Sprintf("%.2f", cpuAllUse)+"  CT:"+fmt.Sprintf("%.2f", fan.CpuTemperatureRead()), img, 0, i*8)
			i++
		} else {
			fmt.Println("get cpu info error:", err)
		}

		memUsage, err := mem.VirtualMemory()

		if err == nil {
			//godump.Dump(memUsage)
			memStr := "M:" + fmt.Sprintf("%.1f", float32(memUsage.Used)/(1024*1024*1024)) + "/" + fmt.Sprintf("%.1f", float32(memUsage.Total)/(1024*1024*1024)) + "G"

			diskStr := ""
			diskUsage, err := disk.Usage("/")
			if nil == err {
				diskStr = "D:" + fmt.Sprintf("%.1f", float32(diskUsage.Used)/(1024*1024*1024)) + "/" + fmt.Sprintf("%.1f", float32(diskUsage.Total)/(1024*1024*1024)) + "G"

				//memDiskStr += " D:" + utilFile.FormatSize(int64(diskUsage.Used)) + "/" + utilFile.FormatSize(int64(diskUsage.Total))
			} else {
				fmt.Println("get disk info error:", err)
			}

			if len(memStr+diskStr) <= 20 {
				textToImage(memStr+" "+diskStr, img, 0, i*8)
				i++
			} else {
				textToImage(memStr, img, 0, i*8)
				i++
				/*textToImage(diskStr, img, 0, i*8)
				i++*/
			}

		} else {
			fmt.Println("get memory info error:", err)
		}

		for {
			if i > 3 {
				break
			}
			viewCacheChanIsEmpty := false
			select {
			case viewCache := <-viewCacheChan:
				if "" != viewCache {
					textToImage(viewCache, img, 0, i*8)
					i++
				}
			default:
				viewCacheChanIsEmpty = true
				break
			}

			if viewCacheChanIsEmpty {
				break
			}

		}

		if i < 4 {
			ips := utils.GetSelfV4Ips()
			for name, ip := range ips {
				if i < 4 {
					textToImage(name+":"+ip, img, 0, i*8)
				} else {
					viewCacheChan <- name + ":" + ip
				}

				i++
			}
		}

		oled.ShowImage(img)
		oled.Display()

	}

	for {
		work()
		time.Sleep(time.Duration(5) * time.Second)
	}

}

func textToImage(text string, img *image.RGBA, x int, y int) {

	ctx := smallfont.Context{
		Font:  smallfont.Font6x8,
		Dst:   img,
		Color: color.White,
	}
	err := ctx.Draw([]byte(text), x, y)
	if err != nil {
		fmt.Println(text, x, y, err)
	}
}
