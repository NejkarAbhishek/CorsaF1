package scheduler

import (
	"f1insight/internal/service"
	"github.com/robfig/cron/v3"
	"log"
)

func StartDataSync() {
	c := cron.New()
	c.AddFunc("@hourly", func() {
		log.Println("Cron job: syncing driver standings...")
		_, err := service.FetchDriverStandings()
		if err != nil {
			log.Println("Error during sync:", err)
		} else {
			log.Println("Driver standings synced successfully.")
		}
	})
	c.Start()
}
