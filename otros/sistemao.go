package otros

import (
	"fmt"
	"runtime"
)

func DatosSistemaOperativos() {
	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("IF Esto es", os)
	} else {
		fmt.Println("ELSE Esto es", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("switch Esto es", os)
	case "darwin":
		fmt.Println("switch Esto es", os)
	default:
		fmt.Printf("switch %s \n", os)
	}
}
