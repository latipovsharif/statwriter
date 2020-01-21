package reports

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
