package main

import (
	"fmt"
	"os"
	"strconv"
)

func CMDLineApp()  {
	args := os.Args[1:]

	if len(args) == 1 && args[0] == "help" {
		fmt.Println("usage: <meal amount> <tip percent>")
		fmt.Println("ex: ./go-std-lib 20 10")

	} else if len(args) != 2 {
		fmt.Println("pleae enter 2 args or checkout the help by typing help")
	} else {
		mealOnly := args[0]
		tipAmt := args[1]

		total := calculateTotalCost(mealOnly, tipAmt)

		fmt.Printf("your meal total is: %.2f \n", total)
	}
}

func calculateTotalCost(mealCost, tipAmt string) float64{

	mealCostF, err := strconv.ParseFloat(mealCost, 64)

	if err != nil {
		fmt.Errorf("error converting %v", mealCost)
	}

	tipAmtF, err := strconv.ParseFloat(tipAmt, 64)

	if err != nil {
		fmt.Errorf("error converting %v", tipAmt)
	}


	return mealCostF + (mealCostF * (tipAmtF/100))
}