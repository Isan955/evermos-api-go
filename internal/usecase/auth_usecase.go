package usecase

import (
	"errors"
	"os"
	"time"

	"evermos-api/internal/entity"
	"evermos-api/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(name, email, phone, password string) error
	Login(email, password string) (string, error)
}

type authUsecase struct {
	userRepo  repository.UserRepository
	storeRepo repository.StoreRepository
}

func NewAuthUsecase(ur repository.UserRepository, sr repository.StoreRepository) AuthUsecase {
	return &authUsecase{ur, sr}
}

func (u *authUsecase) Register(name, email, phone, password string) error {
	if _, err := u.userRepo.FindByEmail(email); err == nil {
		return errors.New("email sudah digunakan")
	}
	if _, err := u.userRepo.FindByPhone(phone); err == nil {
		return errors.New("nomor telepon sudah digunakan")
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &entity.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: string(hashed),
	}
	if err := u.userRepo.Create(user); err != nil {
		return err
	}
	store := &entity.Store{
		Name:   "Toko " + user.Name,
		UserID: user.ID,
	}
	if err := u.storeRepo.Create(store); err != nil {
		return errors.New("gagal membuat toko otomatis")
	}
	return nil
}

func (u *authUsecase) Login(email, password string) (string, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email tidak ditemukan")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("password salah")
	}
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))
	s, err := token.SignedString(secret)
	return s, err
}
