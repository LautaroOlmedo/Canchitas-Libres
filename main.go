package main

import (
	"canchitas-libres/pkg/user/domain/field"
	"fmt"
)

func main() {

	/*	var users []user.User
		user := user.NewUser("lautaro", "olmedo", 1, time.Now(), "lautaroolmedo", "123", "admin")

		users = append(users, user)
		fmt.Println(len(users))*/

	var fields []field.Field
 	field, err := field.NewField("Cancha 7", 2, 6, 40000)
 	if err != nil {
		fmt.Println("Error al crear el campo:", err)
		return
	}
	fields = append(fields, *field)
	fmt.Println(field.FieldType)
}
