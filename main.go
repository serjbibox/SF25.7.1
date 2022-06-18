package main

import (
	"fmt"
	"log"
)

func main() {
	var s string
	fmt.Print("Введите данные в консоль: ")
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели данные: %s\n", s)
}
