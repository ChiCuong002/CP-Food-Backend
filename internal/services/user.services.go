package services

import "food-recipes-backend/internal/repo"

type IUserService interface {
	Register(email string, password string) int
}

type userService struct {
	userRepo repo.IUserRepository
	//...
}

func NewUserService(
	userRepo repo.IUserRepository,
) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, password string) int {
	return 0
}