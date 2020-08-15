package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	BoilingC      Celsius = 100.0
	FreezingC     Celsius = 0.0
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g˚C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g˚F", f)
}
