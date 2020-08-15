package tempconv

// CtoF converts Celsius temperature to Fahrenheit temperature.
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit((c * 5 / 9) + 32)
}

// FtoC converts Fahrenheit temperature to Celsius temperature.
func FtoC(f Fahrenheit) Celsius {
	return Celsius(((f - 32) * 5) / 9)
}
