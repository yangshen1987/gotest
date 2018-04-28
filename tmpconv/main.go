package tmpconv

import (
	"fmt"
	"flag"
)

type Celsius float64
type Fahrenheit float64
type celsiusFalg struct {
	Celsius
} 
const  (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CTof(c Celsius) Fahrenheit  {
	return Fahrenheit(c*9/5+32)
}
func FToc(f Fahrenheit) Celsius  {
	return Celsius((f - 32) *5/9)
}
func (f *celsiusFalg)Set(s string) error  {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", ".C":
		f.Celsius = Celsius(value)
		return nil
	case "F",".F":
		f.Celsius = FToc(Fahrenheit(value))
		return nil

	}
	return fmt.Errorf("invaild temperature %q",s)
}
func CelsiusFlag(name string, value Celsius, usage string) *Celsius  {
	f := celsiusFalg{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
