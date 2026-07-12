package auth

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
)

// Users mewakili data pengguna dengan role
type Users struct {
	Fullname string
	Username string
	Password string
	Role     string // "admin" atau "kasir"
}

// ----- Repository Interface (untuk menyimpan data user) -----
type UserRepository interface {
	Save(user *Users) error
	FindByUsername(username string) (*Users, error)
	FindAll() ([]*Users, error)
	Update(user *Users) error
	Delete(username string) error
}

// ----- Implementasi in-memory (slice) -----
type dataUser struct {
	DB []*Users
}

func NewDataUser() *dataUser {
	return &dataUser{DB: []*Users{}}
}

func (r *dataUser) Save(user *Users) error {
	// Cek duplikat username
	for _, u := range r.DB {
		if u.Username == user.Username {
			return errors.New("username sudah terdaftar")
		}
	}
	r.DB = append(r.DB, user)
	return nil
}

func (r *dataUser) FindByUsername(username string) (*Users, error) {
	for _, u := range r.DB {
		if u.Username == username {
			return u, nil
		}
	}
	return nil, errors.New("user tidak ditemukan")
}

func (r *dataUser) FindAll() ([]*Users, error) {
	return r.DB, nil
}

func (r *dataUser) Update(user *Users) error {
	for i, u := range r.DB {
		if u.Username == user.Username {
			r.DB[i] = user
			return nil
		}
	}
	return errors.New("user tidak ditemukan")
}

func (r *dataUser) Delete(username string) error {
	for i, u := range r.DB {
		if u.Username == username {
			r.DB = append(r.DB[:i], r.DB[i+1:]...)
			return nil
		}
	}
	return errors.New("user tidak ditemukan")
}

// ----- AuthService (menggunakan repository) -----
type AuthService struct {
	repo UserRepository
}

func NewAuthService(repo UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

// hash password ke MD5
func hashPasswd(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}

// Register menambahkan user baru dengan role
func (s *AuthService) Register(fullname, username, password, confirm, role string) error {
	if password != confirm {
		return errors.New("password harus sama dengan confirm password")
	}
	if role != "admin" && role != "kasir" {
		return errors.New("role harus 'admin' atau 'kasir'")
	}

	// Cek apakah username sudah ada (repository akan mengecek juga, tapi kita cek dulu)
	existing, _ := s.repo.FindByUsername(username)
	if existing != nil {
		return errors.New("username sudah terdaftar")
	}

	hashed := hashPasswd(password)
	user := &Users{
		Fullname: fullname,
		Username: username,
		Password: hashed,
		Role:     role,
	}

	if err := s.repo.Save(user); err != nil {
		return err
	}
	fmt.Printf("Registrasi sukses untuk %s (%s)\n", fullname, role)
	return nil
}

// Login memverifikasi kredensial dan mengembalikan data user
func (s *AuthService) Login(username, password string) (*Users, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("username atau password salah")
	}
	if user.Password != hashPasswd(password) {
		return nil, errors.New("username atau password salah")
	}
	fmt.Printf("Selamat datang %s! (role: %s)\n", user.Fullname, user.Role)
	return user, nil
}

// ChangePasswd mengganti password user yang sedang login (berdasarkan username)
func (s *AuthService) ChangePasswd(username, oldPass, newPass, confirm string) error {
	if newPass != confirm {
		return errors.New("password baru harus sama dengan confirm password")
	}

	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	if user.Password != hashPasswd(oldPass) {
		return errors.New("password lama salah")
	}

	user.Password = hashPasswd(newPass)
	if err := s.repo.Update(user); err != nil {
		return err
	}
	fmt.Printf("Password %s berhasil diubah\n", user.Fullname)
	return nil
}
// UpdateUserProfile mengubah fullname dan/atau role (hanya admin)
func (s *AuthService) UpdateUserProfile(username, newFullname, newRole string) error {
    user, err := s.repo.FindByUsername(username)
    if err != nil {
        return err
    }
    if newFullname != "" {
        user.Fullname = newFullname
    }
    if newRole != "" {
        if newRole != "admin" && newRole != "kasir" {
            return errors.New("role harus 'admin' atau 'kasir'")
        }
        user.Role = newRole
    }
    return s.repo.Update(user)
}

// ResetPassword mereset password user tanpa perlu password lama (hanya admin)
func (s *AuthService) ResetPassword(username, newPass string) error {
    user, err := s.repo.FindByUsername(username)
    if err != nil {
        return err
    }
    user.Password = hashPasswd(newPass)
    return s.repo.Update(user)
}

// SearchUsers mencari user berdasarkan username atau fullname (case-insensitive)
func (s *AuthService) SearchUsers(keyword string) ([]*Users, error) {
    all, err := s.repo.FindAll()
    if err != nil {
        return nil, err
    }
    var result []*Users
    keywordLower := strings.ToLower(keyword)
    for _, u := range all {
        if strings.Contains(strings.ToLower(u.Username), keywordLower) ||
            strings.Contains(strings.ToLower(u.Fullname), keywordLower) {
            result = append(result, u)
        }
    }
    return result, nil
}

// (Opsional) Hapus user – hanya untuk admin
func (s *AuthService) DeleteUser(username string) error {
	return s.repo.Delete(username)
}

// (Opsional) Daftar semua user – hanya untuk admin
func (s *AuthService) ListUsers() ([]*Users, error) {
	return s.repo.FindAll()
}