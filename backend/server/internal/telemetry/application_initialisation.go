package command

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	if os.Getenv("ENV") == "DEV" {
		log.SetReportCaller(true)
	}
	log.Printf("Logger initialized")
}
