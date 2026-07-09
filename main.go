package main

import (
	"fmt"
	"project-golang/feature"
	"project-golang/menu"
	"project-golang/utils"
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
	utils.CallClear()
	totalHarga := 0
	var carts []feature.Cart
	var choice string
	filMenu := []menu.User{}
	menuChoice := menu.User{}
choiceKategori:
	for {
		welcMess()
		fmt.Print("Pilih angka (1-5): ")
		fmt.Scanln(&choice)

		filMenu = feature.FilterMenu(menu.ListMenu(), choice)
		if choice == "4" {
				goto cartDisplay
		}
		if filMenu != nil {
			break
		}
	}

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

inputQty:
	for {
		fmt.Print("Input Jumlah Pesanan: ")
		fmt.Scanln(&choice)
		qty, err := strconv.Atoi(choice)
		if err != nil {
			fmt.Println("Error: Input tidak valid! Masukkan Jumlah berupa angka saja.")
			continue
		}
		carts = feature.CartProcess(menuChoice, qty, carts)
		if carts != nil {

			for {
				fmt.Print("Ingin Pesan Lagi Y / N: ")
				fmt.Scanln(&choice)

				if choice == "Y" || choice == "y" {
					utils.CallClear()
					goto choiceKategori
				} else if choice == "N" || choice == "n" {
					break inputQty
				} else {
					fmt.Println("Input Salah, Pilih Y atau N")
					continue
				}
			}

		}

	}
	utils.CallClear()

cartDisplay:
	for {
		if len(carts) == 0 {
			fmt.Println("Keranjang Masih Kosong Silahkan Pesan dulu")
			goto choiceKategori
		}

		totalHarga += feature.DisplayCart(carts)
		fmt.Print("Ingin Pesan Lagi Y / N: ")
		fmt.Scanln(&choice)

		if choice == "Y" || choice == "y" {
			utils.CallClear()
			goto choiceKategori
		} else if choice == "N" || choice == "n" {
			break
		} else {
			fmt.Println("Input Salah, Pilih Y atau N")
			continue
		}

	}

}
