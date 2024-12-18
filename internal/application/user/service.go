package user

import "user/internal/domain/user"

type UserService struct {
	repo user.UserRepository
}

func NewUserService(repo user.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) error {
	user := user.NewUser(user.NewUserID(), name, email)
	return s.repo.Create(user)
}