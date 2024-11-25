package main

import (
	"fmt"
	// "price-concurrency/cmdmanager"
	"price-concurrency/filemanager"

	"price-concurrency/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		// cmdm := cmdmanager.New()
		fm := filemanager.New("price.txt", fmt.Sprintf("result_%.0f.json", taxRate))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

  for index, _ := range taxRates {
    select {
    case err := <-errorChans[index]:
      if err != nil {
        fmt.Println(err)
      }
    case <-doneChans[index]:
      fmt.Println("Done!")
    }
  }

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
