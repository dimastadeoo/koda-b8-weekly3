package menu

import (
	"encoding/json"
	"strconv"
)

type User struct {
	IdMenu      int    `json:"idmenu"`
	Name        string `json:"nama"`
	Price       int    `json:"harga"`
	Description string `json:"deskripsi"`
	Category    string `json:"kategori"`
}



func ListMenu(jsonMenu []byte) []User {
	// file, err := os.ReadFile("menu.json")
	// if err != nil {
	// 	panic("Error Message: " + err.Error())
	// }
	var ListMenu []User

	err := json.Unmarshal(jsonMenu, &ListMenu)
	if err != nil {
		panic("Error Message: " + err.Error())

	}
	return ListMenu
}

func FilterMenu(items []User, targetCategory string) []User {
	var result []User
	for _, item := range items {
		if item.Category == targetCategory {
			result = append(result, item)
		} else if id, _ := strconv.Atoi(targetCategory); item.IdMenu == id {
			result = append(result, item)
		}
	}
	return result
}
