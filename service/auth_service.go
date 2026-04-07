package service

import (
	models "asset-management/model"
	"asset-management/repository"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

var jwtSecret = []byte("RahasiaNegara123")

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GenerateToken tetap di sini karena ini adalah logika bisnis keamanan
func (s *UserService) GenerateToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Login sekarang berfungsi sebagai gerbang utama Admin
func (s *UserService) Login(email, password string) (*models.User, string, error) {
	// 1. Cari user berdasarkan email
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// 2. Bandingkan password bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// 3. Generate Token langsung di sini agar Handler lebih bersih
	token, err := s.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, "", errors.New("failed to generate session")
	}

	return user, token, nil
}

func (s *UserService) ListUsers(search string, limit, offset int) ([]models.User, int64, error) {
	return s.repo.GetAll(search, limit, offset)
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

// UpdateUser: Admin bisa ganti password sendiri atau ganti email
func (s *UserService) UpdateUser(u *models.User, rawPassword string) error {
	if rawPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return s.repo.Update(u)
}
