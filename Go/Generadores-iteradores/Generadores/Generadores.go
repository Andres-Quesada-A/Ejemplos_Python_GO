package main

//En Go no existen los generadores, pero se pueden crear funciones y permitan obtener valores de forma asincrónica
import (
	"fmt"
)

func main() {
	evenNumbers := generateEvenNumbers()

	for n := range evenNumbers {
		fmt.Println(n)
	}
}

func generateEvenNumbers() <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i <= 10; i++ {
			if i%2 == 0 {
				ch <- i
			}
		}
	}()

	return ch
}
