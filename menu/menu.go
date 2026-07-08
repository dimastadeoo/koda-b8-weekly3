package menu

import (
	"encoding/json"
	"log"
	"os"
)

type Pengguna struct {
	IdMenu    int    `json:"id_menu"`
	Nama      string `json:"nama"`
	Harga     int    `json:"harga"`
	Deskripsi string `json:"deskripsi"`
	Kategori  string `json:"kategori"`
}


func ListMenu() []Pengguna{
	file, err := os.Open("menu.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var daftarMenu []Pengguna

	// Gunakan decoder langsung dari file
	err = json.NewDecoder(file).Decode(&daftarMenu)
	if err != nil {
		log.Fatal(err)
	}
	return daftarMenu
}
