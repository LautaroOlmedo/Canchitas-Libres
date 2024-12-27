package storage

type Slice struct {
}

type MySQL struct{}

type Mongo struct {
	// deberia estar el atributo para hablar con la base de datos
}

func NewSlice() *Slice {
	return &Slice{
		// deberia estar el atributo para hablar con la base de datos
	}
}

func NewMySQL() *MySQL {
	return &MySQL{}
}

func NewMongo() *Mongo {
	return &Mongo{}
}

// hacer implementacion de postgres
