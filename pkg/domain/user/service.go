package user

// patron repository. Creo una interfaz que puede ser implementada por muchos "adaptadores". (MySQL, MOngoDB, Redis, etc...)
type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) (*User, error)
	// agregar UpdateUser
	DeleteByID(id int) error
}

type UserService struct {
	UserRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{UserRepository: userRepository}
}
