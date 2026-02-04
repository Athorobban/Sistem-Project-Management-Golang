package services

import (
	"errors"

	"github.com/Athorobban/Sistem-Project-Management-Golang/models"
	"github.com/Athorobban/Sistem-Project-Management-Golang/repositories"
	"github.com/Athorobban/Sistem-Project-Management-Golang/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Register(user *models.User) error
	Login(email,password string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	GetByPublicID(id string) (*models.User, error)
	GetAllPagination(filter,sort string, limit,offset int)([]models.User,int64, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type userService struct{
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService)Register(user *models.User) error {
	existingUser, _ := s.repo.FindByEmail(user.Email)
	if existingUser.InternalID !=0 {
		return errors.New("email already registered")
	}
	hased,err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hased
	user.Role = "user"
	user.PublicID = uuid.New()
	return s.repo.Create(user)
}

func (s *userService) Login(email,password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil,errors.New("Invalid Credentials")
	}
	if !utils.CheckPasswordHash(password,user.Password){
		return nil, errors.New("Invalid Credentials")
	}
	return user, nil
}

func (s *userService) GetByID(id uint) (*models.User, error){
	return s.repo.FindByID(id)
}
func (s *userService) GetByPublicID(id string) (*models.User, error){
	return s.repo.FindByPublicID(id)
}
func (s * userService) GetAllPagination(filter,sort string, limit,offset int)([]models.User,int64, error)  {
	return s.repo.FindAllPagination(filter,sort,limit,offset)
}
func (s *userService) Update(user *models.User) error {
	return s.repo.Update(user)
}
func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}