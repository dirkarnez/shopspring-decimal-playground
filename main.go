package main

import (
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

func main() {
	fmt.Println("Why don't you just use float64?")
	var n float64 = 0
	for i := 0; i < 1000; i++ {
		n += .01
	}
	fmt.Println("> 1000 * 0.01 expected 10.0, but get", n)
	fmt.Println()

	fmt.Println("Why don't you just use big.Rat?")
	z, _ := new(big.Rat).SetString("1")     // z = 1
	three, _ := new(big.Rat).SetString("3") // three = 3
	x := new(big.Rat).Quo(z, three)         // x = z / 3, which is 1/3
	y := new(big.Rat).Quo(z, three)         // y = z / 3, which is also 1/3
	z = z.Sub(z, x)                         // z -= x
	z = z.Sub(z, y)                         // z -= y, now z should also be 1/3
	s := new(big.Rat).Add(x, y)             // s = x + y, which is 1/3 + 1/3 = 2/3
	s.Add(s, z)                             // s += z

	// where did the other 0.001 go?
	fmt.Printf("> %s + %s + %s = %s", x.FloatString(3), y.FloatString(3), z.FloatString(3), s.FloatString(3))
	fmt.Println()
	fmt.Println()

	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}
