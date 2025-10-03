package main

/*
#include <stdlib.h>
#include <math.h>
*/
import "C"

import "math"

func Circle_Square(D C.double) C.double {
	return math.Pi * D * D / 4.0
}

//export Permeability_Cylinder
func Permeability_Cylinder(L, D, mu, dP, Qk C.double,
	Km, Kmkm, KmD *C.double) C.int {

	if L > 0 && D > 0 && mu > 0 && dP != 0 && Qk > 0 {
		F := C.M_PI * D * D / 4.0
		*Km = (L * mu * Qk) / (dP * F)
		*Kmkm = *Km * C.pow(10, 15)
		*KmD = *Kmkm / 0.9869
		return 0
	}

	*Km = 0
	*Kmkm = 0
	*KmD = 0
	return 1
}

func main() {
	print("Hello")
}
