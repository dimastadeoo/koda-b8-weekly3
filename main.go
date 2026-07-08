package main

import (
	"fmt"
	"project-golang/menu"
)

func main() {
	

	// Cetak hasil
	for i, item := range menu.ListMenu() {
		fmt.Printf("%d: %+v\n", i, item)
	}
}