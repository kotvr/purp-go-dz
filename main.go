package main

import (
	"fmt"
)

const usdInEUR = 0.88
const usdInRUB = 78.75

func main() {
	fmt.Println("Конвертер валют")
	userOrig, userValue, userTarget := getUserInput()
	switch {
	case userOrig == "USD" && userTarget == "EUR":
		fmt.Print("Из ", userOrig, " в ", userTarget, " выходит: ", userValue*usdInEUR, " ", userTarget)
	case userOrig == "USD" && userTarget == "RUB":
		fmt.Print("Из ", userOrig, " в ", userTarget, " выходит: ", userValue*usdInRUB, " ", userTarget)
	case userOrig == "EUR" && userTarget == "USD":
		fmt.Print("Из ", userOrig, " в ", userTarget, " выходит: ", userValue/usdInEUR, " ", userTarget)
	case userOrig == "EUR" && userTarget == "RUB":
		fmt.Print("Из ", userOrig, " в ", userTarget, " выходит: ", userValue*getEURInRUB(), " ", userTarget)
	case userOrig == "RUB" && userTarget == "USD":
		fmt.Print("Из ", userOrig, " в ", userTarget, " выходит: ", userValue/usdInRUB, " ", userTarget)
	default:
		fmt.Print("Из ", userOrig, " в ", userTarget, " выходит: ", userValue/getEURInRUB(), " ", userTarget)
	}
}

func getEURInRUB() float64 {
	var eurInRUB float64 = usdInRUB / usdInEUR
	return eurInRUB
}

func getUserInput() (string, float64, string) {
	var userValue float64
	var userOrig string
	var userTarget string
	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&userOrig)
	userOrig = checkNameOrigValutes(userOrig)
	fmt.Print("Введите количество валюты: ")
	fmt.Scan(&userValue)
	userValue = checkValue(userValue)
	switch {
	case userOrig == "USD":
		fmt.Print("Введите целевую валюту (EUR, RUB): ")
	case userOrig == "EUR":
		fmt.Print("Введите целевую валюту (USD, RUB): ")
	default:
		fmt.Print("Введите целевую валюту (EUR, USD): ")
	}
	fmt.Scan(&userTarget)
	userTarget = checkNameTargetValutes(userOrig, userTarget)
	return userOrig, userValue, userTarget
}

func checkNameOrigValutes(userValutes string) string {
	for {
		if userValutes == "USD" || userValutes == "EUR" || userValutes == "RUB" {
			return userValutes
		} else {
			fmt.Println("Невернная валюта, повторите ввод: ")
			fmt.Scan(&userValutes)
		}
	}
}

func checkNameTargetValutes(userOrigValutes string, userTargetValutes string) string {
	for {
		if userTargetValutes == "USD" || userTargetValutes == "EUR" || userTargetValutes == "RUB" && userTargetValutes != userOrigValutes {
			return userTargetValutes
		} else {
			fmt.Println("Невернная валюта, повторите ввод: ")
			fmt.Scan(&userTargetValutes)
		}
	}
}

func checkValue(userValue float64) float64 {
	for {
		if userValue > 0 {
			return userValue
		} else {
			fmt.Println("Невернное количество валюты, повторите ввод:")
			fmt.Scan(&userValue)
		}
	}
}
