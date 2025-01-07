package main

import (
	"canchitas-libres/pkg/domain/field"
	"canchitas-libres/pkg/repository/field/storage"
	"canchitas-libres/pkg/repository/rest"
	"fmt"
)

// "canchitas-libres/pkg/user/domain/field"
// "fmt"

func main() {

	mongoRepository := storage.NewMongo()
	//sliceRepository := storage.NewSlice()
	fieldService := field.NewFieldService(mongoRepository)
	_, err := fieldService.GetField()
	if err != nil {
		fmt.Println(err)
	}

	// handler := rest.NewHandler()
	// server := rest.NewServer(handler)
	// server.Start()

	h := rest.NewHandlerField()
	server := rest.NewServer(h)
	server.Start()


	// var fields []field.Field
 	// field, err := field.NewField("Cancha 7", 2, 6, 40000)
 	// if err != nil {
	// 	fmt.Println("Error al crear el campo:", err)
	// 	return
	// }
	// fields = append(fields, *field)
	// fmt.Println(field.FieldType)
}
