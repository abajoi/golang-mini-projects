package main

import (
	"entities"
	"fmt"
)

func main() {
	product1 := entities.Product(nil)

	fmt.Println("Product 1 Info")
	fmt.Println(product1.ToString())
	fmt.Println("Total: ", product1.Total())

	product2 := entities.NewProduct("p02", "name 2", 2, 7, true)
	fmt.Println("Product 2 Info")
	fmt.Println(product2.ToString())
	fmt.Println("Total: ", product2.Total())

}
