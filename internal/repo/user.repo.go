package repo

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {

}

// GetUserByEmail implements IUserRepository
func (us *userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}