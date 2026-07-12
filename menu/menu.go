package menu

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Menu struct {
	IdMenu      int    `json:"idmenu"`
	Name        string `json:"nama"`
	Price       int    `json:"harga"`
	Description string `json:"deskripsi"`
	Category    string `json:"kategori"`
}

func ListMenu(jsonMenu []byte) []Menu {
	// file, err := os.ReadFile("menu.json")
	// if err != nil {
	// 	panic(err.Error())
	// }
	var ListMenu []Menu

	err := json.Unmarshal(jsonMenu, &ListMenu)
	if err != nil {
		panic(err.Error())

	}
	return ListMenu
}

func FilterMenu(items []Menu, target string) []Menu {
	var result []Menu
	for _, item := range items {
		if item.Category == target {
			result = append(result, item)
		} else if id, _ := strconv.Atoi(target); item.IdMenu == id {
			result = append(result, item)
		} else if strings.Contains(strings.ToLower(item.Name), target) {

			result = append(result, item)
		}
	}
	return result
}
