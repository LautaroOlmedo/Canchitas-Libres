package user

func (s *UserService) GetAll() ([]User, error) {
	return s.UserRepository.GetAll() // no llamo a mi persistencia directamente, si no que llamo al metodo que implementa la interfaz
}
