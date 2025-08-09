package user

type UserService1 interface {
	CreateUser(expression, pass string) (*User, error)
	GetAllUser() ([]User, error)
	GetUser(id int) (User, error)
	UpdataUser(id int, user User) (User, error)
	DeleteUser(id int) error
	GetUserForTasks(userID int) (User, error)
}

type CUsersServive struct {
	repo UsersRepository
}

// GetUserForTasks implements UserService1.
func (c *CUsersServive) GetUserForTasks(userID int) (User, error) {
	return c.repo.GetUserForTasksByRepo(userID)
}

func NewUserService(r UsersRepository) UserService1 {
	return &CUsersServive{repo: r}

}

// CreateUser implements UserService.
func (c *CUsersServive) CreateUser(expression string, pass string) (*User, error) {
	ur := &User{Email: expression, Password: pass}
	if err := c.repo.CreateUser(ur); err != nil {
		return nil, err
	}
	return ur, nil
}

// DeleteUser implements UserService.
func (c *CUsersServive) DeleteUser(id int) error {
	return c.repo.DeleteUser(id)
}

// GetAllUser implements UserService.
func (c *CUsersServive) GetAllUser() ([]User, error) {
	return c.repo.GetAllUser()
}

// GetUser implements UserService.
func (c *CUsersServive) GetUser(id int) (User, error) {
	return c.repo.GetUser(id)
}

// UpdataUser implements UserService.
func (c *CUsersServive) UpdataUser(id int, user User) (User, error) {
	us, err := c.repo.GetUser(id)
	if err != nil {
		return User{}, err
	}

	if user.Email != "" {
		us.Email = user.Email
	}
	if user.Password != "" {
		us.Password = user.Password
	}

	if err := c.repo.UpdataUser(us); err != nil {
		return User{}, err
	}
	return us, nil

}
