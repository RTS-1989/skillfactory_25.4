package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите данные: ")
	sc.Scan()
	txt := sc.Text()
	fmt.Printf("Вы ввели следующие данные: %v\n", txt)
}
