package service

import "ecom/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetUser(id string) (string, string) {
	return us.userRepo.GetUser(id)
}
