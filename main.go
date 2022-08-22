package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите целое число: ")
	sc.Scan()
	txt := sc.Text()
	fmt.Printf("Вы ввели число: %v\n", txt)
}
