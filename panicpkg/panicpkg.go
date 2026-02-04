package panicpkg

//panic : erro do programador/erro execução tempo

import "fmt"

func main() {
	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	panic("Pânico")

}
