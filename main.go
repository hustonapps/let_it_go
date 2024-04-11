package main

import (
	"fmt"
	"let_it_go/cmdmanager"
	// "let_it_go/filemanager"
	"let_it_go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("could not process job: ", err)
		}
	}
}