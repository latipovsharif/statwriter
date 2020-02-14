package reports

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type devices int

// devices list
const (
	UNKNOWN devices = 0
	PC      devices = 1
	TAB     devices = 2
	MOBILE  devices = 3
	CONSOLE devices = 4
)

// writeDeviceStat get and write device stat to file
func writeDeviceStat(db *sqlx.DB, d devices, filename string) error {
	log.Info("retrieving device stats")
	gs, err := deviceStat(db, d)
	if err != nil {
		return errors.Wrap(err, "cannot get device stat")
	}

	log.Info("device stats retrieved")
	log.Info("writing device stats")

	if err := write(gs, filename); err != nil {
		return errors.Wrap(err, "cannot write device stat")
	}

	log.Info("finished writing device stats")

	return nil
}

// deviceStat get stat grouped by devices
func deviceStat(db *sqlx.DB, d devices) ([]BaseStr, error) {
	var dpc []BaseStr

	err := db.Select(&dpc, `SELECT group_field, req_count, resp_count, buy_count, click_count,
			if (buy_count = 0, 0, toFloat64(round(click_count / buy_count * 100, 3))) as ctr,
			if (buy_count = 0, 0, toFloat64(round(buyout_sum / buy_count * 1000, 3))) AS cpm,
			if (buy_count = 0, 0, toFloat64(round(first_sum / buy_count * 1000, 3))) AS first_price,
			if (click_count = 0, 0 , toFloat64(round(buyout_sum / click_count, 2))) AS click_price_ssp,
			if (click_count = 0, 0, toFloat64(round(click_sum / click_count, 2))) as click_price_dsp,
			round(buyout_sum, 2) as ssp_flow,
			round(click_sum, 2) as dsp_flow,
			round((dsp_flow - ssp_flow), 2) as profit
		FROM (
			SELECT req_device_device_type AS group_field , count(*) as req_count
			FROM bid_requests
			WHERE req_site_name != '' AND req_ssp_id IN (2) AND toDate(req_date_time) = (toDate(now()) - 1) and group_field = ?
			GROUP BY group_field 
		)
		ALL LEFT JOIN (
			SELECT group_field , resp_count, buy_count, click_count, buyout_sum, click_sum, first_sum
			FROM (
				SELECT req_device_device_type AS group_field, count(*) as resp_count
				FROM bid_responses
				WHERE req_site_name != '' AND req_ssp_id IN (2) AND toDate(resp_date_time) = (toDate(now()) - 1)  and group_field = ?
				GROUP BY group_field 
			)
			ALL LEFT JOIN (
				SELECT group_field , buy_count, click_count, buyout_sum, click_sum, first_sum
				FROM (
					SELECT req_device_device_type AS group_field, count(*) as buy_count, sum(buyout_price) as buyout_sum, sum(resp_seat_bid_bid_price) as first_sum
					FROM buyouts
					WHERE req_site_name != '' AND req_ssp_id IN (2) AND toDate(buyout_date_time) = (toDate(now()) - 1) and group_field = ?
					GROUP BY group_field
				)
				ALL LEFT JOIN (
					SELECT req_device_device_type AS group_field, count(*) as click_count, sum(click_price) as click_sum
					FROM clicks
					WHERE req_site_name != '' AND req_ssp_id IN (2) AND toDate(buyout_date_time) = (toDate(now()) - 1) and group_field = ?
					GROUP BY group_field
				)
				USING(group_field)
			)
			USING(group_field)
		)
		USING(group_field);`, d, d, d, d)
	return dpc, err
}
