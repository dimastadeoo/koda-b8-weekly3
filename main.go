package main

import (
	_ "embed"
	"fmt"
	"log"
	"project-golang/auth"
	"project-golang/feature"
	"project-golang/menu"
	"project-golang/utils"
	"strconv"
)

//go:embed menu.json
var jsonMenu []byte

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
	totalPrice := 0
	var carts []feature.Cart
	var choice string
	filMenu := []menu.User{}
	menuChoice := menu.User{}

	defer func (){
		if r := recover(); r != nil {
			fmt.Println("Error: ", r)
		}
	}()

	// Inisialisasi repository dan service
	repo := auth.NewDataUser()
	service := auth.NewAuthService(repo)

	// Cek apakah sudah ada admin, jika belum buat admin default
	adminExists := false
	users, _ := service.ListUsers()
	for _, u := range users {
		if u.Role == "admin" {
			adminExists = true
			break
		}
	}
	if !adminExists {
		fmt.Println("Belum ada admin. Silakan buat admin awal.")
		// Minta input admin pertama
		fullname := utils.ReadString("Nama Lengkap Admin: ")
		username := utils.ReadString("Username Admin: ")
		password := utils.ReadString("Password Admin: ")
		confirm := utils.ReadString("Konfirmasi Password: ")
		err := service.Register(fullname, username, password, confirm, "admin")
		if err != nil {
			log.Fatal("Gagal membuat admin:", err)
		}
		fmt.Println("Admin berhasil dibuat! Silakan login.")
		utils.PressEnter("Tekan enter untuk lanjut")
	}
authentikasi:
	for {
		utils.CallClear()
		fmt.Println("========= SISTEM LOGIN =========")
		fmt.Println("1. Login sebagai Admin")
		fmt.Println("2. Login sebagai Kasir")
		fmt.Println("3. Keluar")
		choice := utils.ReadString("Pilih: ")

		switch choice {
		case "1":
			username := utils.ReadString("Username: ")
			password := utils.ReadString("Password: ")
			user, err := service.Login(username, password)
			if err != nil {
				fmt.Println("Login gagal:", err)
				utils.PressEnter("Tekan Enter untuk lanjut")
				continue
			}
			if user.Role != "admin" {
				fmt.Println("Akses ditolak. Bukan admin.")
				utils.PressEnter("Tekan Enter untuk lanjut")

				continue
			}
			fmt.Printf("Selamat datang Admin %s!\n", user.Fullname)
			utils.PressEnter("Tekan Enter untuk lanjut")
			auth.AdminMenu(service, user)

		case "2":
			username := utils.ReadString("Username: ")
			password := utils.ReadString("Password: ")
			user, err := service.Login(username, password)
			if err != nil {
				fmt.Println("Login gagal:", err)
				utils.PressEnter("Tekan Enter untuk lanjut")
				continue
			}
			if user.Role != "kasir" {
				fmt.Println("Akses ditolak. Bukan kasir.")
				utils.PressEnter("Tekan Enter untuk lanjut")
				continue
			}
			fmt.Printf("Selamat datang Kasir %s!\n", user.Fullname)
			utils.PressEnter("Tekan Enter untuk lanjut")
			// Langsung ke menu pemesanan (ganti dengan fungsi Anda)
			break authentikasi
			// Setelah selesai, kembali ke login

		case "3":
			fmt.Println("Terima kasih, sampai jumpa!")
			return

		default:
			fmt.Println("Pilih 1-3")
			utils.PressEnter("Tekan Enter untuk lanjut")
		}
	
	}


choiceKategori:
	for {
		welcMess()
		fmt.Print("Pilih angka (1-5): ")
		fmt.Scanln(&choice)

		filMenu = feature.FilterMenu(menu.ListMenu(jsonMenu), choice)
		if choice == "4" {
				goto cartDisplay
		}
		if filMenu != nil {
			break
		}
	}
inputMenu:
	for {
		fmt.Print("Pilih Nomor Menu / Ketik 0 untuk kembali ke awal: ")
		fmt.Scanln(&choice)
		if choice == "0" {
			utils.CallClear()
			goto choiceKategori
		}
		menuChoice = feature.DetailMenu(filMenu, choice)
		if (menuChoice != menu.User{}) {
			break
		}
	}

inputQty:
	for {
		fmt.Print("Input Jumlah Pesanan / input 0 untuk kembali ke menu sebelumnya: ")
		fmt.Scanln(&choice)
		qty, err := strconv.Atoi(choice)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		} else if qty == 0 {
			fmt.Println("Silahkan pilih Nomor Menu kembali")
			goto inputMenu
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

		totalPrice = feature.DisplayCart(carts)
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

	for {
		fmt.Print("Input Pembayaran: ")
		fmt.Scanln(&choice)
		payment, err := strconv.Atoi(choice)

		if err != nil {
			fmt.Println("Error: Input tidak valid! Masukkan hanya angka")
			continue
		}else if payment < totalPrice {
			fmt.Println("Error: Pembayaran Kurang! Input Pembayaran harus lebih dari sama dengan Total Pembayaran")
			continue
		}

		carts = feature.Payment(payment, totalPrice, carts)

		for {
			fmt.Print("Ingin Memulai Pesanan Lagi Y / N: ")
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
		
		break
	}

}
