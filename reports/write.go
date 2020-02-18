package reports

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/tealeg/xlsx"
)

// write write data to filename
func write(baseStr []BaseStr, filename string) error {
	if len(baseStr) == 0 {
		return nil
	}

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var err error

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		file = xlsx.NewFile()
		sheet, err = file.AddSheet("Sheet0")
		if err != nil {
			return errors.Wrap(err, "cannot create sheet")
		}
	} else {
		file, err = xlsx.OpenFile(filename)
		if err != nil {
			return errors.Wrapf(err, "cannot open file: %v", filename)
		}
		sheet = file.Sheets[0]
	}

	for _, item := range baseStr {
		row := sheet.AddRow()
		row.AddCell().Value = time.Now().Format("2006-01-02")
		row.AddCell().Value = item.GroupField
		row.AddCell().Value = fmt.Sprintf("%d", item.RequestCount)
		row.AddCell().Value = fmt.Sprintf("%d", item.ResponseCount)
		row.AddCell().Value = fmt.Sprintf("%d", item.BuyCount)
		row.AddCell().Value = fmt.Sprintf("%d", item.ClickCount)
		row.AddCell().Value = fmt.Sprintf("%f", item.CTR)
		row.AddCell().Value = fmt.Sprintf("%f", item.CPM)
		row.AddCell().Value = fmt.Sprintf("%f", item.FirstPrice)
		row.AddCell().Value = fmt.Sprintf("%f", item.ClickPriceSSP)
		row.AddCell().Value = fmt.Sprintf("%f", item.ClickPriceDSP)
		row.AddCell().Value = fmt.Sprintf("%f", item.FlowSSP)
		row.AddCell().Value = fmt.Sprintf("%f", item.FlowDSP)
		row.AddCell().Value = fmt.Sprintf("%f", item.Profit)
	}

	err = file.Save(filename)
	if err != nil {
		return errors.Wrap(err, "cannot save file")
	}

	return nil
}
