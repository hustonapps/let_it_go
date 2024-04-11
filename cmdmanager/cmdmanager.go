package cmdmanager

import "fmt"

type CMDManager struct {}

func (cm CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter prices, one per line. Enter 'done' when finished.")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "done" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (cm CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
