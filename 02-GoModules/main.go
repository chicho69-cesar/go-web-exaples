package main

import (
	"fmt"
	"github.com/chicho69-cesar/go-web-exaples/02-gomodules/maths"
)

func main() {
	fmt.Println("Hello World!!!")

	fmt.Println("La suma de 5 + 4 = ", maths.Sum(5, 4))
	fmt.Println("La resta de 15 - 3 = ", maths.Subs(15, 3))
	fmt.Println("La multiplicación de 45 * 2 = ", maths.Mult(45, 2))
	fmt.Println("La división de 21 / 7 = ", maths.Div(21, 7))
	fmt.Println("La potencia de 6 ^ 3 = ", maths.Pot(6, 3))
}
