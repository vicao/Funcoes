// Closure : função "chamar" uma variável que está em outra função
package closurepkg

import "fmt"

func Closure() {
	x := 0

	numero := func() int {
		x++
		return x
	}

	fmt.Println(numero())

}
