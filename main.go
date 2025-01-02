package main

import (
	"canchitas-libres/pkg/repository/rest"
)

func main() {

	/*mongoRepository := storage.NewMongo()
	//sliceRepository := storage.NewSlice()
	userService := user.NewUserService(mongoRepository)
	_, err := userService.GetUser()
	if err != nil {
		fmt.Println(err)
	}*/

	handler := rest.NewHandler()
	server := rest.NewServer(handler)
	server.Start()

}
