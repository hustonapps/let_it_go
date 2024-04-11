package main

import (
	"fmt"
	// "let_it_go/cmdmanager"
	"let_it_go/filemanager"
	"let_it_go/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate * 100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index], index)
	}

	for index := range taxRates {
		select {
		case err := <- errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <- doneChans[index]:
			fmt.Println(("done"))
		}
	}
}