package main

import (
	"TheGoLanguage/Chapter2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, val := range os.Args[1:] {
		t, err := strconv.ParseFloat(val, 64)
		if err != nil {
			_ = fmt.Errorf("Error %v", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t)
		f := tempconv.Fahrenheit(t)

		fmt.Printf("%s = %s and %s = %s\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))

	}
}
