package menu

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type User struct {
	IdMenu    int    `json:"idmenu"`
	Name      string `json:"nama"`
	Price     int    `json:"harga"`
	Description string `json:"deskripsi"`
	Category  string `json:"kategori"`
}

func ListMenu() []User{
	file, err := os.ReadFile("menu.json")
	if err != nil {
		log.Fatal(err)
	}
	var ListMenu []User
	
	err = json.Unmarshal(file, &ListMenu)
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
        }else if id, _:= strconv.Atoi(targetCategory); item.IdMenu == id{
            result = append(result, item)
		}
    }
    return result
}
