package handler

import (
	"fmt"
	"time"

	// log "github.com/sirupsen/logrus"
	"github.com/robfig/cron/v3"
)
// enable_scheduler: true
// count down : "@every 25s"
// time on day : "TZ=Asia/Bangkok 59 23 * * *" #time_zone MM:HH
func StartScheduler() {
	_enableScheduler := false
	if _enableScheduler {
		fmt.Println("Start Scheduler")
		startCheckInBroadcastScheduler()
		startAlertScheduler()
	}
}

func startCheckInBroadcastScheduler() {
	_timeScheduler := "TZ=Asia/Bangkok 11 14 * * *" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, broadcastJob)
	c.Start()
	msg := fmt.Sprintf("Schedule Boibot Send CheckIn Broadcast running at: %s", c.Entries()[0].Next.String())
	fmt.Println("StartCheckInBroadcastScheduler : ", msg)
}

func startAlertScheduler() {
	_timeScheduler := "@every 10s" //time_zone MM:HH
	c := cron.New()
	c.AddFunc(_timeScheduler, alertCheckInJob)
	c.Start()
	msg := fmt.Sprintf("Schedule Alert finish check in step running at: %s", c.Entries()[0].Next.String())
	fmt.Println("StartAlertScheduler : ", msg)
}


func broadcastJob() {
	fmt.Println("===============================")
	fmt.Println("Check In Broadcast Scheduler : ", time.Now())
	fmt.Println("===============================")

}

func alertCheckInJob() {
	fmt.Println("##############################")
	fmt.Println("Alert Check In Scheduler : ", time.Now())
	fmt.Println("##############################")

}
