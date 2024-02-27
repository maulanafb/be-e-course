package user

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserByID(ID int) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(input FormUpdateUserInput) (User, error)
	CheckingEmail(email string) (User, error)
	FindOrCreateUserByEmail(email string, name string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	// user.Role = input.Role

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.Image = fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID ")
	}

	return user, nil
}

func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}

func (s *service) UpdateUser(input FormUpdateUserInput) (User, error) {
	user, err := s.repository.FindByID(input.ID)
	if err != nil {
		return user, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = input.Role

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) CheckingEmail(email string) (User, error) {
	check, err := s.repository.FindByEmail(email)
	fmt.Println(check)
	if err != nil {
		return check, err
	}
	return check, nil
}

func (s *service) FindOrCreateUserByEmail(email string, userName string) (User, error) {
	// 1. Cari user berdasarkan email
	user := User{}

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) { // Periksa apakah error adalah karena user tidak ditemukan
			return user, err // return error jika errornya bukan karena user tidak ditemukan
		}
	}
	randomPassword, err := generateRandomPassword(12)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(randomPassword), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Name = userName
	user.Email = email
	user.Role = "user"
	// 2. Jika tidak ditemukan, buat user baru
	if user.ID == 0 {

		user, err = s.repository.Save(user)
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func generateRandomPassword(length int) (string, error) {
	randomBytes := make([]byte, (length*3)/4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:length], nil
}
