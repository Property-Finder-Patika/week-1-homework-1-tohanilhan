package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Below functions are related with Exercise 2.1 on the page 42 of the GoPL book

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit(k*9/5 - 459.67)
}
