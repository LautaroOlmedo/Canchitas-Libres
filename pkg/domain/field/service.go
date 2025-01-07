package field

type FieldRepository interface {
	GetAll() ([]Field, error)
	GetByFieldID(id int) (*Field, error)
	DeleteByFieldID(id int) error
}

type FieldService struct {
	FieldRepository FieldRepository
}

func NewFieldService(FieldRepository FieldRepository) *FieldService {
	return &FieldService{FieldRepository: FieldRepository}
}