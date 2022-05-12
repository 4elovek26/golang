package main

import (
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	"os"  // пакет используется для проверки ответа, не удаляйте его
)

func printError(value interface{}) {
	fmt.Printf("value=%v: %T", value, value)
	os.Exit(0)
}

func operat(value1, value2, oper interface{}) {
	switch oper {
	case "+":
		result := value1.(float64) + value2.(float64)
		fmt.Printf("%.4f", result)
	case "-":
		result := value1.(float64) - value2.(float64)
		fmt.Printf("%.4f", result)
	case "*":
		result := value1.(float64) * value2.(float64)
		fmt.Printf("%.4f", result)
	case "/":
		result := value1.(float64) / value2.(float64)
		fmt.Printf("%.4f", result)
	default:
		fmt.Println("неизвестная операция")

	}
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса

	switch value1.(type) {
	case float64:
		switch value2.(type) {
		case float64:
			operat(value1, value2, operation)

		default:
			printError(value2)
		}
	default:
		printError(value1)
	}

}
