package main

import (
	"fmt"
	"math"
)
// the z² − x is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. This general approach is called Newton's method.
func Sqrt(x float64) float64 {
	z := float64(x)/2
//	z := float64(1)
	for i:=0;i<10;i++{
	    z -= (z*z-x)/(2*z)
		fmt.Printf("%v ", z)
		if d:=math.Abs(math.Sqrt(x)-z); d < 1.0e-16 {
			fmt.Println("\ni=%v d=%v",i,d)
		    return z
		}
	}
	fmt.Println("\n",math.Sqrt(x)-z)
	return z
}

func main() {
	fmt.Println(Sqrt(100))
}

//output:
//26 14.923076923076923 10.812053925455988 10.030495203889796 10.000046356507898 10.000000000107446 10 
//i=%v d=%v 6 0
//10