package main

import (
	"canchitas-libres/pkg/domain/reserve"
	"fmt"
	"time"
)

func main() {

	// mongoRepository := storage.NewMongo()
	// //sliceRepository := storage.NewSlice()
	// userService := user.NewUserService(mongoRepository)
	// _, err := userService.GetUser()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	var reserves []reserve.Reserve

	reserve1, err := reserve.NewReserve("iddelususaio123", "iddelacanchas5623", time.Now(), 120)
	if err != nil {
		fmt.Println("Error al crear la reserva:", err)
		return
	}
	reserves = append(reserves, *reserve1)

	reserve2, err := reserve.NewReserve("iddelususaio123", "iddelacanchas5623", time.Now(), 120)
	if err != nil {
		fmt.Println("Error al crear la reserva:", err)
		return
	}
	reserves = append(reserves, *reserve2)

	fmt.Println(len(reserves))
}
