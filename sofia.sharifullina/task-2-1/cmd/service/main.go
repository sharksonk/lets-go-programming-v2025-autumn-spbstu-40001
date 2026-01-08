package main

import "fmt"

const (
	minTemperature = 15
	maxTemperature = 30
)

type ConditionerT struct {
	minTemperature int
	maxTemperature int
}

func (cond *ConditionerT) changeTemperature(sign string, degrees int) {
	switch sign {
	case ">=":
		if degrees >= cond.minTemperature {
			cond.minTemperature = degrees
		}

	case "<=":
		if degrees <= cond.maxTemperature {
			cond.maxTemperature = degrees
		}
	}
}

func main() {
	var departNum int

	_, err := fmt.Scan(&departNum)
	if err != nil {
		fmt.Println("Invalid input", err)

		return
	}

	for range departNum {
		var emplCount int

		_, err := fmt.Scan(&emplCount)
		if err != nil {
			fmt.Println("Invalid input", err)

			return
		}

		conditioner := ConditionerT{minTemperature, maxTemperature}

		for range emplCount {
			var sign string

			_, err = fmt.Scan(&sign)
			if err != nil {
				fmt.Println("Invalid input", err)

				return
			}

			var degrees int

			_, err = fmt.Scan(&degrees)
			if err != nil {
				fmt.Println("Invalid input", err)

				return
			}

			conditioner.changeTemperature(sign, degrees)

			if conditioner.minTemperature <= conditioner.maxTemperature {
				fmt.Println(conditioner.minTemperature)
			} else {
				fmt.Println("-1")
			}
		}
	}
}
