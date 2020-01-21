package main

import (
	"dailystatuploader/reports"
	"fmt"
	"os"
	"path"
	"sync"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := getConfigs(); err != nil {
		panic(fmt.Sprintf("cannot read configs: %v", err))
	}

	logPath := viper.GetString("logPath")
	if logPath == "" {
		logPath = "."
	}

	f, err := os.OpenFile(path.Join(logPath, "log.log"), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(fmt.Sprintf("cannot open log file to write %v", err))
	}

	log.SetOutput(f)

	doJob()
}

func doJob() {
	// Wait group to wait forever otherwise application will close immediately after start

	log.Info("starting job")
	cronString := viper.GetString("cronString")
	if cronString == "" {
		cronString = "@midnight"
	}

	log.Infof("cron value is %v", cronString)

	c := cron.New()
	c.AddFunc(cronString, reports.WriteStats)
	c.Start()

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}

func getConfigs() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "cannot read configs")
	}
	return nil
}
