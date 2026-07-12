package auth

import (
	"fmt"
	"project-golang/utils"
)

func AdminMenu(s *AuthService, admin *Users) {
    utils.CallClear()
    for {
        fmt.Println("==================== MENU ADMIN ====================")
        fmt.Println("1. Tambah User (admin/kasir)")
        fmt.Println("2. Lihat Semua User")
        fmt.Println("3. Update Data User")
        fmt.Println("4. Hapus User")
        fmt.Println("5. Cari User")
        fmt.Println("6. Logout")
        choice := utils.ReadString("Pilih: ")

        switch choice {
        case "1":
            // Tambah user
            fullname := utils.ReadString("Nama Lengkap: ")
            username := utils.ReadString("Username: ")
            pass := utils.ReadString("Password: ")
            confirm := utils.ReadString("Konfirmasi Password: ")
            role := ""
            for {
                role = utils.ReadString("Role (admin/kasir): ")
                if role == "admin" || role == "kasir" {
                    break
                }
                fmt.Println("Role harus 'admin' atau 'kasir'")
            }
            err := s.Register(fullname, username, pass, confirm, role)
            if err != nil {
                fmt.Println("Gagal tambah user:", err)
            } else {
                fmt.Println("User berhasil ditambahkan.")
            }
            utils.PressEnter("Tekan Enter untuk Kembali ke menu awal")

        case "2":
            // Lihat semua user
            users, err := s.ListUsers()
            if err != nil {
                fmt.Println("Error:", err)
            } else {
                fmt.Printf("\n%-5s | %-15s | %-20s | %-10s\n", "NO", "USERNAME", "NAMA LENGKAP", "ROLE")
                fmt.Println("------------------------------------------------------------")
                for i, u := range users {
                    fmt.Printf("%-5d | %-15s | %-20s | %-10s\n", i+1, u.Username, u.Fullname, u.Role)
                }
            }
            utils.PressEnter("Tekan Enter untuk Kembali ke menu awal")

        case "3":
            // Update user
            targetUsername := utils.ReadString("Masukkan username yang akan di-update: ")
            // Cek apakah user ada
            target, err := s.ListUsers() // atau pakai FindByUsername
            if err != nil {
                fmt.Println("Error:", err)
                utils.PressEnter("Tekan enter untuk lanjut")
                continue
            }
            var found *Users
            for _, u := range target {
                if u.Username == targetUsername {
                    found = u
                    break
                }
            }
            if found == nil {
                fmt.Println("User tidak ditemukan.")
                utils.PressEnter("Tekan Enter untuk lanjut")
                continue
            }

            fmt.Printf("Data saat ini: Nama=%s, Role=%s\n", found.Fullname, found.Role)
            fmt.Println("Pilih yang ingin diubah:")
            fmt.Println("1. Nama Lengkap")
            fmt.Println("2. Password (reset)")
            fmt.Println("3. Role")
            fmt.Println("4. Kembali")
            subChoice := utils.ReadString("Pilih: ")

            switch subChoice {
            case "1":
                newName := utils.ReadString("Nama baru: ")
                if err := s.UpdateUserProfile(targetUsername, newName, ""); err != nil {
                    fmt.Println("Gagal update:", err)
                } else {
                    fmt.Println("Nama berhasil diupdate.")
                }
            case "2":
                newPass := utils.ReadString("Password baru: ")
                confirmPass := utils.ReadString("Konfirmasi password baru: ")
                if newPass != confirmPass {
                    fmt.Println("Password tidak cocok.")
                } else {
                    if err := s.ResetPassword(targetUsername, newPass); err != nil {
                        fmt.Println("Gagal reset password:", err)
                    } else {
                        fmt.Println("Password berhasil direset.")
                    }
                }
            case "3":
                newRole := ""
                for {
                    newRole = utils.ReadString("Role baru (admin/kasir): ")
                    if newRole == "admin" || newRole == "kasir" {
                        break
                    }
                    fmt.Println("Role harus 'admin' atau 'kasir'")
                }
                if err := s.UpdateUserProfile(targetUsername, "", newRole); err != nil {
                    fmt.Println("Gagal update role:", err)
                } else {
                    fmt.Println("Role berhasil diupdate.")
                }
            case "4":
                // kembali
            default:
                fmt.Println("Pilihan tidak valid.")
            }
            utils.PressEnter("Tekan Enter untuk Kembali ke menu awal")

        case "4":
            // Hapus user (tidak boleh hapus diri sendiri)
            targetUsername := utils.ReadString("Masukkan username yang akan dihapus: ")
            if targetUsername == admin.Username {
                fmt.Println("❌ Anda tidak dapat menghapus akun sendiri!")
                utils.PressEnter("Tekan enter untuk lanjut")
                continue
            }
            // Konfirmasi
            confirm := utils.ReadString("Yakin ingin menghapus user " + targetUsername + "? (y/n): ")
            if confirm != "y" && confirm != "Y" {
                fmt.Println("Penghapusan dibatalkan.")
                utils.PressEnter("Tekan enter untuk lanjut")
                continue
            }
            if err := s.DeleteUser(targetUsername); err != nil {
                fmt.Println("Gagal hapus user:", err)
            } else {
                fmt.Println("User berhasil dihapus.")
            }
            utils.PressEnter("Tekan Eneter untuk lanjut")

        case "5":
            // Cari user
            keyword := utils.ReadString("Masukkan kata kunci (username atau nama): ")
            results, err := s.SearchUsers(keyword)
            if err != nil {
                fmt.Println("Error:", err)
            } else if len(results) == 0 {
                fmt.Println("Tidak ada user yang cocok.")
            } else {
                fmt.Printf("\n%-5s | %-15s | %-20s | %-10s\n", "NO", "USERNAME", "NAMA LENGKAP", "ROLE")
                fmt.Println("------------------------------------------------------------")
                for i, u := range results {
                    fmt.Printf("%-5d | %-15s | %-20s | %-10s\n", i+1, u.Username, u.Fullname, u.Role)
                }
            }
                        utils.PressEnter("Tekan Eneter untuk lanjut")

        case "6":
            fmt.Println("Logout dari admin...")
            utils.PressEnter("Tekan Eneter untuk lanjut")

            return

        default:
            fmt.Println("Pilih hanya angka 1 - 6")
            utils.PressEnter("Tekan Eneter untuk lanjut")
        }
        utils.CallClear()
    }
}