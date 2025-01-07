package field

func (f *FieldService) GetField() ([]Field, error) {
	return f.FieldRepository.GetAll()
}