// this file converts two fahrenheit temperatures to Celsius
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gºF = %gºC \n", freezingF, FtoC(freezingF))
	fmt.Printf("%gºF = %gºC \n", boilingF, FtoC(boilingF))
}

func FtoC(f float32) float32 {
	return ((f - 32) * 5) / 9
}
