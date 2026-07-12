package feature

import (
	"fmt"
	"project-golang/auth"
	"project-golang/utils"
)

func AdminLogin(service *auth.AuthService) {
	username := utils.ReadString("Username: ")
	password := utils.ReadString("Password: ")
	user, err := service.Login(username, password)
	if err != nil {
		fmt.Println("Login gagal:", err)
		utils.PressEnter("Tekan Enter untuk lanjut")
		return
	}
	if user.Role != "admin" {
		fmt.Println("Akses ditolak. Bukan admin.")
		utils.PressEnter("Tekan Enter untuk lanjut")
		return
	}
	fmt.Printf("Selamat datang Admin %s!\n", user.Fullname)
	utils.PressEnter("Tekan Enter untuk lanjut")
	auth.AdminMenu(service, user)
}

func KasirLogin(service *auth.AuthService) {
	username := utils.ReadString("Username: ")
	password := utils.ReadString("Password: ")
	user, err := service.Login(username, password)
	if err != nil {
		fmt.Println("Login gagal:", err)
		utils.PressEnter("Tekan Enter untuk lanjut")
		return
	}
	if user.Role != "kasir" {
		fmt.Println("Akses ditolak. Bukan kasir.")
		utils.PressEnter("Tekan Enter untuk lanjut")
		return
	}
	fmt.Printf("Selamat datang Kasir %s!\n", user.Fullname)
	utils.PressEnter("Tekan Enter untuk lanjut")
}
