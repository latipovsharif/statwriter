package main

import (
	"dailystatuploader/reports"
	"dailystatuploader/writer"
	"fmt"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

func main() {
	connect, err := sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?password=123")
	if err != nil {
		log.Fatal(err)
	}

	gs, err := reports.DateStat(connect)
	if err != nil {
		fmt.Println(err)
	}

	if err := writer.Write(gs, "/home/user/Desktop/date_stat.xlsx"); err != nil {
		fmt.Println(err)
	}
}
