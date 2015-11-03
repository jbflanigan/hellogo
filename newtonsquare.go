package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0;
	for i := 0; i < 10 ; i++ {
	    z = z - ((z*z-x)/(2*z))
	}
    return z
}

func SqrtStop(x float64) (float64, int) {
	z := 1.0;
	delta := 1.0;
	i := 0
	const MinDelta = 0.000000001;
	for delta > MinDelta {
		newZ := z - ((z*z-x)/(2*z))
		delta = math.Abs(z - newZ);
		z = newZ
		i++;
	}
    return z, i
}


func main() {
	for i := 1; i < 101 ; i++ {
		f := float64(i);
		stop, count := SqrtStop(f);
		fmt.Printf("square root of %d. Newton 10: %f Newton delta (in %d) %f math.Sqrt %f\n", i, Sqrt(f), count, stop, math.Sqrt(f))
	}
}
