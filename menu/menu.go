package menu

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	IdMenu    int    `json:"idmenu"`
	Name      string `json:"nama"`
	Price     int    `json:"harga"`
	Description string `json:"deskripsi"`
	Category  string `json:"kategori"`
}


func ListMenu() []User{
	file, err := os.Open("menu.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ListMenu []User

	// Gunakan decoder langsung dari file
	err = json.NewDecoder(file).Decode(&ListMenu)
	if err != nil {
		log.Fatal(err)
	}
	return ListMenu
}

func FilterMenu(items []User, targetCategory string) []User {
    var result []User
    for _, item := range items {
        if item.Category == targetCategory {
            result = append(result, item)
        }
    }
    return result
}
