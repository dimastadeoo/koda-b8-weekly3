package main

import (
	"fmt"
	"project-golang/feature"
	"project-golang/menu"
)
var welcMess = func(){
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
	welcMess()
	var choice string
	fmt.Print("Pilih angka (1-5): ")
	fmt.Scanln(&choice)

	feature.FilterMenu(menu.ListMenu(), choice)


}