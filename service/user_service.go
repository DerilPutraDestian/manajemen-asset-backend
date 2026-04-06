package service

import (
	models "asset-management/model"
	"asset-management/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(u *models.User, rawPassword string) error {
	// 1. Validasi: Cek apakah email sudah terdaftar
	existing, _ := s.repo.GetByEmail(u.Email)
	if existing != nil && existing.ID != "" {
		return errors.New("email already registered")
	}

	// 2. Hashing Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return s.repo.Create(u)
}

func (s *UserService) ListUsers(search string, limit, offset int) ([]models.User, int64, error) {
	return s.repo.GetAll(search, limit, offset)
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) UpdateUser(u *models.User) error {
	return s.repo.Update(u)
}

func (s *UserService) DeleteUser(u *models.User) error {
	return s.repo.Delete(u)
}
func (s *UserService) Login(email, password string) (*models.User, error) {
	// 1. Cari user berdasarkan email melalui repository
	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// 2. Bandingkan password input dengan password hash di database
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// Jangan beri tahu secara spesifik apakah password atau email yang salah (keamanan)
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
