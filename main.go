package main

import (
	"canchitas-libres/pkg/domain/user"
	"canchitas-libres/pkg/repository/user/storage"
	"fmt"
)

func main() {

	mongoRepository := storage.NewMongo()
	//sliceRepository := storage.NewSlice()
	userService := user.NewUserService(mongoRepository)
	_, err := userService.GetUser()
	if err != nil {
		fmt.Println(err)
	}

}
