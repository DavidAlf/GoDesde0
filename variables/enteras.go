package variables

import (
	"fmt"
)

func MuestroEnteros() {
	var intcomun int

	intde32 := int32(10)
	intde64 := int64(100)

	intcomun = 1
	fmt.Println("IntComun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}
