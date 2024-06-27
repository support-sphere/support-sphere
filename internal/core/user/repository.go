package user

type Repository interface {
	Create(user *User) error
	FindByUsername(username string) (*User, error)
}
