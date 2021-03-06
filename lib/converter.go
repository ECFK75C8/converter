package main

import (
	"fmt"
	"math"
)

type Converter struct{}

type (
	Second      float64
	Millisecond float64
	Minute      float64
	Feet        float64
	Centimeter  float64
	Celsius     float64
	Fahrenheit  float64
	Radian      float64
	Degree      float64
	Kilogram    float64
	Pounds      float64
)

func (cvr Converter) SecondToMillisecond(s Second) Millisecond {
	return Millisecond(s * 1000)
}

func (cvr Converter) MillisecondToSecond(m Millisecond) Second {
	return Second(m / 1000)
}

func (cvr Converter) SecondToMinute(s Second) Minute {
	return Minute(s * 60)
}

func (cvr Converter) MinuteToSecond(m Minute) Second {
	return Second(m / 60)
}

func (cvr Converter) CentimeterToFeet(c Centimeter) Feet {
	return Feet(c * 30.48)
}

func (cvr Converter) FeetToCentimeter(f Feet) Centimeter {
	return Centimeter(f / 30.48)
}

func (cvr Converter) CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (cvr Converter) FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (cvr Converter) DegreeToRadian(d Degree) Radian {
	return Radian(d * (math.Pi / 180))
}

func (cvr Converter) RadianToDegree(r Radian) Degree {
	return Degree(r * (180 / math.Pi))
}

func (cvr Converter) KilogramToPounds(k Kilogram) Pounds {
	return Pounds(k * 2.2)
}

func (cvr Converter) PoundsToKilogram(p Pounds) Kilogram {
	return Kilogram(p / 2.2)
}

func main() {
	cvr := Converter{}
	fmt.Println(cvr.CelsiusToFahrenheit(32))
	fmt.Println(cvr.FahrenheitToCelsius(120.56))
	fmt.Println(cvr.SecondToMinute(2))
	fmt.Println(cvr.SecondToMillisecond(1))
	fmt.Println(cvr.KilogramToPounds(50))
	fmt.Println(cvr.DegreeToRadian(30))
	fmt.Println(cvr.CentimeterToFeet(15))
}
