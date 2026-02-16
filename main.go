package main

import "fmt"

func main() {
	myBill := newBill("Mwencha's bill")

	fmt.Println(myBill.format())

}
