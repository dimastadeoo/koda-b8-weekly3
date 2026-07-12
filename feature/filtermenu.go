package feature

import (
	"fmt"
	"project-golang/menu"
	"project-golang/utils"
)

func FilterMenu(dataMenu []menu.Menu, category string) []menu.Menu {
	var filMenu []menu.Menu

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
		choice := ""
		for {
			utils.CallClear()
			fmt.Println("--------------------------------------------------------------")
			choice = utils.ReadString("Cari Menu berdasarkan Kategori atau nama menu: ")
			filMenu = menu.FilterMenu(dataMenu, choice)
			if len(filMenu) == 0{
				fmt.Println("Hasil Pencarian Menu Tidak ada")
				utils.PressEnter("Tekan Enter untuk ulangi lagi")
				return nil
			}else{
				break
			}
		}
		fmt.Printf("---------------------Pencarian (%s)-----------------------------\n", choice)
	case "5":
		utils.CallClear()
		return nil
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
