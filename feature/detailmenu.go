package feature

import (
	"fmt"
	"project-golang/menu"
)

func DetailMenu(filMenu []menu.Menu, choice string) menu.Menu {
	menuChoice := menu.FilterMenu(filMenu, choice)
	if menuChoice == nil {
		fmt.Println("Menu tidak ada, Silahkan Pilih lagi!")
		return menu.Menu{}
	}

	fmt.Println(`--------------------------------------------------------------`)
	fmt.Println(`---------------------Detail Pesanan---------------------------`)

	for _, menu := range menuChoice {
		fmt.Printf("Nama: %s\n", menu.Name)
		fmt.Printf("Harga: %s\n", FormatRupiah(menu.Price))
		fmt.Printf("Deskripsi: %s\n", menu.Description)
	}
	fmt.Println(`--------------------------------------------------------------`)

	return menuChoice[0]
}
