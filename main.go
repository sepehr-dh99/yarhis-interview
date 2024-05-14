package main

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

type IUserService interface {
	GetUser(ID *uint) (*User, error)
}

type UserService struct {
	userRepo IMainRepository[User]
}

func NewServiceService() IUserService {
	db := db.DB

	return &UserService{
		userRepo: NewMainRepository[User](db),
	}
}

func (us *UserService) GetUser(ID *uint) (*User, error) {
	user, err := us.userRepo.Get(ID)
	if err != nil {
		return nil, err
	}

	return &User{}, err
}
