package main

import (
	"dailystatuploader/reports"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/robfig/cron/v3"
)

func main() {
	doJob()
}

func doJob() {
	c := cron.New()
	c.AddFunc("@midnight", reports.WriteStats)
	c.Start()
}
