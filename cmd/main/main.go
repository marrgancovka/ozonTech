package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	app "ozonTech/internal/app"
	"time"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.DateTime,
		FullTimestamp:   true,
	})
	application := app.NewApp(logger)
	if err := application.Run(); err != nil {
		logger.Fatal(err)
	}

}
