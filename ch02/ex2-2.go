package main

import (
	"flag"
	"fmt"
)

type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

const (
	meterPerFeet     = 0.3048
	kilogramPerPound = 0.45359237
)

var f = flag.Float64("f", -1, "feet")
var m = flag.Float64("m", -1, "meter")
var p = flag.Float64("p", -1, "pound")
var k = flag.Float64("k", -1, "kilogram")

func main() {
	flag.Parse()
	if *f >= 0.0 {
		_f := Feet(*f)
		fmt.Printf("%v == %v\n", _f, feetToMeter(_f))
	}
	if *m >= 0.0 {
		_m := Meter(*m)
		fmt.Printf("%v == %v\n", _m, meterToFeet(_m))
	}
	if *p >= 0.0 {
		_p := Pound(*p)
		fmt.Printf("%v == %v\n", _p, poundToKilogram(_p))
	}
	if *k >= 0.0 {
		_k := Kilogram(*k)
		fmt.Printf("%v == %v\n", _k, kilogramToPound(_k))
	}
}

func feetToMeter(f Feet) Meter {
	return Meter(f * meterPerFeet)
}

func meterToFeet(m Meter) Feet {
	return Feet(m / meterPerFeet)
}

func poundToKilogram(p Pound) Kilogram {
	return Kilogram(p * kilogramPerPound)
}

func kilogramToPound(k Kilogram) Pound {
	return Pound(k / kilogramPerPound)
}

func (f Feet) String() string {
	return fmt.Sprintf("%gft", f)
}
func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
func (p Pound) String() string {
	return fmt.Sprintf("%glb", p)
}
func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}
