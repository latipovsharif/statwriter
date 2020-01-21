package main

import (
	"dailystatuploader/reports"
	"fmt"
	"sync"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func main() {
	if err := getConfigs(); err != nil {
		panic(fmt.Sprintf("cannot read configs: %v", err))
	}

	doJob()
}

func doJob() {
	// Wait group to wait forever otherwise application will close immediately after start
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

	c := cron.New()
	c.AddFunc("@midnight", reports.WriteStats)
	c.Start()
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
