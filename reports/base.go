package reports

import (
	"log"

	"github.com/jmoiron/sqlx"
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
	db, err := sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?password=123")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	writeDateStat(db, "")
	writeCountryStat(db, "")
	writeDeviceStat(db, PC, "")
	writeDeviceStat(db, MOBILE, "")
	writeSubscription(db, "")
}
