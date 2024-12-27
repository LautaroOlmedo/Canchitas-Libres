package user

func (s *UserService) GetUser() ([]User, error) {
	return s.UserRepository.GetAll() // no llamo a mi persistencia directamente, si no que llamo al metodo que implementa la interfaz
}
