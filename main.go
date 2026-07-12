package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
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
	4. Cari Menu
	5. Lihat Keranjang`)
	fmt.Println(`--------------------------------------------------------------`)
}

func main() {
	utils.CallClear()
	totalPrice := 0
	var carts []feature.Cart
	filMenu := []menu.Menu{}
	menuChoice := menu.Menu{}

	defer func() {
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
	for _, user := range users {
		if user.Role == "admin" {
			adminExists = true
			break
		}
	}
	if !adminExists {
		fmt.Println("Belum ada admin. Silakan buat admin awal.")
		var fullname, username, password, confirm string
		// Minta input admin pertama
		for {
			fullname = utils.ReadString("Nama Lengkap Admin: ")
			username = utils.ReadString("Username Admin: ")
			password = utils.ReadString("Password Admin: ")
			confirm = utils.ReadString("Konfirmasi Password: ")
			if fullname == "" || username == "" || password == "" || confirm == "" {
				fmt.Println("Data pendaftaran tidak boleh ada kosong")
				utils.PressEnter("Tekan enter untuk lanjut")
				utils.CallClear()
				continue
			} else if confirm != password {
				fmt.Println("confirm password tidak sama dengan password")
				utils.PressEnter("Tekan enter untuk lanjut")
				continue
			} else {
				break
			}
		}
		err := service.Register(fullname, username, password, confirm, "admin")
		if err != nil {
			log.Fatal("Gagal membuat admin:", err)
		}
		fmt.Println("Admin berhasil dibuat! Silakan login.")
		utils.PressEnter("Tekan enter untuk lanjut")
		utils.CallClear()
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
			feature.AdminLogin(service)
		case "2":
			feature.KasirLogin(service)
			break authentikasi
		case "3":
			fmt.Println("Terima kasih, sampai jumpa!")
			os.Exit(1)
		default:
			fmt.Println("Pilih 1-3")
			utils.PressEnter("Tekan Enter untuk lanjut")
		}
	}

	utils.CallClear()
choiceKategori:
	for {
		welcMess()
		choice := utils.ReadString("Pilih angka (1-5): ")

		filMenu = feature.FilterMenu(menu.ListMenu(jsonMenu), choice)
		if choice == "5" {
			goto cartDisplay
		}
		if filMenu != nil {
			break
		}
	}
inputMenu:
	for {
		choice := utils.ReadString("Pilih Nomor Menu / Ketik 0 untuk kembali ke awal: ")
		if choice == "0" {
			utils.CallClear()
			goto choiceKategori
		}
		menuChoice = feature.DetailMenu(filMenu, choice)
		if (menuChoice != menu.Menu{}) {
			break
		}
	}

inputQty:
	for {
		choice := utils.ReadString("Input Jumlah Pesanan / input 0 untuk kembali ke menu sebelumnya: ")

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
				choice = utils.ReadString("Ingin Pesan Lagi Y / N: ")

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
		choice := utils.ReadString("Ingin Pesan Lagi Y / N: ")

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
		choice := utils.ReadString("Input Pembayaran: ")

		payment, err := strconv.Atoi(choice)

		if err != nil {
			fmt.Println("Error: Input tidak valid! Masukkan hanya angka")
			continue
		} else if payment < totalPrice {
			fmt.Println("Error: Pembayaran Kurang! Input Pembayaran harus lebih dari sama dengan Total Pembayaran")
			continue
		}

		carts = feature.Payment(payment, totalPrice, carts)

		for {
			choice := utils.ReadString("Ingin Memulai Pesanan Lagi Y / N: ")

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
