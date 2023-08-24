package v1

import (
	"strconv"
	"time"

	"github.com/baaj2109/shorturl/common"
	"github.com/baaj2109/shorturl/dao"
)

var addClick = make(chan string)

const (
	_ = iota
	ASave
	AShutDown
)

// async click count measurde
func Clicker() {
	go func() {
		clicks := make(map[string]int, 100)
		for {
			select {
			case code := <-addClick:
				if code == strconv.Itoa(ASave) {
					go dao.UrlCodeDAO{}.SaveClicks(clicks)
					clicks = make(map[string]int, 100)
				} else if code == strconv.Itoa(AShutDown) {
					go dao.UrlCodeDAO{}.SaveClicks(clicks)
				} else {
					if value, ok := clicks[code]; ok {
						clicks[code] = value + 1
					} else {
						clicks[code] = 1
					}
					if len(clicks) > 1000 {
						go dao.UrlCodeDAO{}.SaveClicks(clicks)
						clicks = make(map[string]int, 100)
					}
				}

			case <-time.After(60 * time.Second):
				addClick <- strconv.Itoa(ASave)
			}
		}
	}()
}

func StopRecord() {
	addClick <- strconv.Itoa(AShutDown)
	common.SugarLogger.Info("The program is going to shutdown,save clicks,waiting for 5s")
	time.Sleep(5 * time.Second)
}
