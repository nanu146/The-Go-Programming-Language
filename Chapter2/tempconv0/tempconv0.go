package tempconv0

type Celsius float32
type Fahrenheit float32

const (
	AbsoluteZeroC Celsius = -273.15
	freezingC     Celsius = 0.0
	boilingC      Celsius = 100.0
)

func main() {

}

func FtoC(f Fahrenheit) Celsius {
	return Celsius(((f - 32) * 5) / 9)
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}
