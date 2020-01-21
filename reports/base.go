package reports

import (
	"fmt"
	"os"
	"path"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// BaseStr struct for writer
type BaseStr struct {
	GroupField    string  `db:"group_field"`
	RequestCount  int32   `db:"req_count"`
	ResponseCount int32   `db:"resp_count"`
	BuyCount      int32   `db:"buy_count"`
	ClickCount    int32   `db:"click_count"`
	CTR           float64 `db:"ctr"`
	CPM           float64 `db:"cpm"`
	FirstPrice    float64 `db:"first_price"`
	ClickPriceSSP float64 `db:"click_price_ssp"`
	ClickPriceDSP float64 `db:"click_price_dsp"`
	FlowSSP       float64 `db:"ssp_flow"`
	FlowDSP       float64 `db:"dsp_flow"`
	Profit        float64 `db:"profit"`
}

// WriteStats write all stats to files
func WriteStats() {
	fmt.Println("writing stats")

	connectionString := viper.GetString("connectionString")
	if connectionString == "" {
		panic("connection string not set")
	}

	db, err := sqlx.Open("clickhouse", connectionString)
	defer db.Close()

	if err != nil {
		panic("cannot open database")
	}

	reportFolderPath := viper.GetString("reportFolderPath")
	if _, err := os.Stat(reportFolderPath); err != nil {
		panic(fmt.Sprintf("report folder path does not exists"))
	}

	if err := writeDateStat(db, path.Join(reportFolderPath, "date_stat.xlsx")); err != nil {
		fmt.Println(err)
	}
	if err := writeCountryStat(db, path.Join(reportFolderPath, "country_stat.xlsx")); err != nil {
		fmt.Println(err)
	}
	if err := writeDeviceStat(db, PC, path.Join(reportFolderPath, "device_pc_stat.xlsx")); err != nil {
		fmt.Println(err)
	}
	if err := writeDeviceStat(db, MOBILE, path.Join(reportFolderPath, "device_mobile_stat.xlsx")); err != nil {
		fmt.Println(err)
	}
	if err := writeSubscription(db, path.Join(reportFolderPath, "subscription_stat.xlsx")); err != nil {
		fmt.Println(err)
	}
}
