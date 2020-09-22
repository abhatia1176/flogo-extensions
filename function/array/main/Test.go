package main

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
	"math/big"
)

func bigsum(v1 float64, v2 float64) float64 {
	var prec uint
	prec = 53
	d := new(big.Float).SetPrec(prec).SetFloat64(v1)
	e := new(big.Float).SetPrec(prec).SetFloat64(v2)
	t := new(big.Float)
	t = t.Add(t, e)
	t = t.Add(t, d)
	fmt.Println("Sum Output with big float:", t)

	fmt.Printf("%T\n", t)
	fmt.Printf("t = %.10g (%s, prec = %d, acc = %s)\n", &t, t.Text('p', 0), t.Prec(), t.Acc())

	sum1 := fmt.Sprintf("%+v", t)
	out, _ := coerce.ToFloat64(sum1)
	fmt.Println("Float64 is: ", out)
	return out
}

func main() {
	fmt.Println("Hello, playground")
	var a, b, c, sum float64
	//var sum float64
	a = -0.3
	b = 2.6
	//a = -0.3234223334562345
	//b = 2.63333333434512343
	c = 3.7
	half := new(big.Float).SetPrec(50).SetFloat64(a + b)
	fmt.Println("Output with float64:", a+b)
	fmt.Println("Output with float64:", a+c)
	fmt.Println("Output with big float:", half)

	d := new(big.Float).SetPrec(52).SetFloat64(a)
	e := new(big.Float).SetPrec(52).SetFloat64(b)
	t := new(big.Float)
	t = t.Add(t, e)
	t = t.Add(t, d)
	fmt.Println("Sum Output with big float:", t)

	fmt.Printf("%T\n", t)
	fmt.Printf("t = %.10g (%s, prec = %d, acc = %s)\n", &t, t.Text('p', 0), t.Prec(), t.Acc())

	sum1 := fmt.Sprintf("%+v", t)
	sum, _ = coerce.ToFloat64(sum1)
	fmt.Println("Float64 is: ", sum)

	for _, test := range []struct {
		val1 float64
		val2 float64 // float32, float64, or string (== 512bit *Float)
		want float64
	}{
		// from fmt/fmt_test.go
		{2.6, -0.3, 2.3},
		{3.7, -0.3, 3.4},
		{2.0, -0.3, 1.7},
	} {
		out := bigsum(test.val1, test.val2)
		fmt.Println("Out from sum function is: ", out)
		if test.want != out {
			fmt.Errorf("got %v (%v); want %v", out, test.want)
		} else {
			fmt.Println("sum is successful")
		}
		fmt.Println("***Float64 is: ", out)
	}

}
