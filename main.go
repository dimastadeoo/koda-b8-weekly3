package main

import (
	"fmt"
	"project-golang/feature"
	"project-golang/menu"
	"strconv"
)

var welcMess = func() {
	fmt.Println(`--------------------Selamat Datang di POPEYE------------------`)
	fmt.Println(`Silahkan Pilih Ingin Pesan Apa:
	1. Paket Makan
	2. Makanan
	3. Minuman
	4. Lihat Keranjang
	5. Keluar`)
	fmt.Println(`--------------------------------------------------------------`)
}

func main() {
	var carts []feature.Cart

	var choice string
	filMenu := []menu.User{}
	choiceKategori:
	for {
		welcMess()
		fmt.Print("Pilih angka (1-5): ")
		fmt.Scanln(&choice)

		filMenu = feature.FilterMenu(menu.ListMenu(), choice)
		if filMenu != nil {
			break
		}
	}

	menuChoice := menu.User{}
	// choiceDetail:
	for {
		fmt.Print("Pilih Nomor Menu / Ketik 0 untuk kembali ke awal: ")
		fmt.Scanln(&choice)
		if choice == "0" {
			goto choiceKategori
		}
		menuChoice = feature.DetailMenu(filMenu, choice)
		if (menuChoice != menu.User{}) {
			break
		}
	}

	for {
		fmt.Print("Input Jumlah Pesanan: ")
		fmt.Scanln(&choice)
		qty, err := strconv.Atoi(choice)
		if err != nil {
			fmt.Println("Error: Input tidak valid! Masukkan Jumlah berupa angka saja.")
			continue
		}
		carts = feature.CartProcess(menuChoice, qty)
		if carts != nil {
			break
		}

	}

}
