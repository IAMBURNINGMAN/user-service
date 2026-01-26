package user

type UserService interface {
	CreateUser(user UserStruct) (UserStruct, error)
	GetAllUsers() ([]UserStruct, error)
	GetUserById(id uint) (UserStruct, error)
	UpdateUser(id uint, user UserStruct) (UserStruct, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo UserRepository
}

func (s *userService) CreateUser(user UserStruct) (UserStruct, error) {
	newUser := UserStruct{
		Email:    user.Email,
		Password: user.Password,
	}
	if err := s.repo.CreateUser(&newUser); err != nil {
		return UserStruct{}, err
	}
	return newUser, nil
}

func (s *userService) GetAllUsers() ([]UserStruct, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserById(id uint) (UserStruct, error) {
	return s.repo.GetUserById(id)
}

func (s *userService) UpdateUser(id uint, user UserStruct) (UserStruct, error) {
	existingUser, err := s.repo.GetUserById(id)
	if err != nil {
		return UserStruct{}, err
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if err := s.repo.UpdateUser(&existingUser); err != nil {
		return UserStruct{}, err
	}
	return existingUser, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}
