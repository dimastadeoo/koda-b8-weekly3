package feature

import (
	"fmt"
	"os"
	"project-golang/menu"
	"project-golang/utils"
	"time"
)

func FilterMenu(dataMenu []menu.User, category string) []menu.User {
	var filMenu []menu.User

	switch category {
	case "1":
		utils.CallClear()
		filMenu = menu.FilterMenu(dataMenu, "paket")
		fmt.Println("---------------------Menu Paket-------------------------------")
	case "2":
		utils.CallClear()
		filMenu = menu.FilterMenu(dataMenu, "makanan")
		fmt.Println("---------------------Menu Makanan-----------------------------")
	case "3":
		utils.CallClear()
		filMenu = menu.FilterMenu(dataMenu, "minuman")
		fmt.Println("---------------------Menu Minuman-----------------------------")
	case "4":
		utils.CallClear()
		fmt.Println("Masuk Ke Cart...")
		time.Sleep(2*time.Second)
		return nil
	case "5":
		os.Exit(1)
	default:
		utils.CallClear()
		fmt.Println("Pilihan tidak valid!")
		return nil
	}

	for _, m := range filMenu {
		fmt.Printf("No. %d %s - (%s)\n", m.IdMenu, m.Name, FormatRupiah(m.Price))
	}

	return filMenu

}
