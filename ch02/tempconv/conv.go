package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celcius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - absoluteZeroC)
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

func KToC(k Kelvin) Celcius {
	return Celcius(k + absoluteZeroK)
}

func KtoF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}
