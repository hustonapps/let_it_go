package prices

import (
	"fmt"
	"let_it_go/conversion"
	"let_it_go/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]string
	IOManager filemanager.FileManager
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return;
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return;
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	err := job.IOManager.WriteResult(job)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate: taxRate,
		IOManager: fm,
	}
}