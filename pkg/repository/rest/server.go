package rest

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	//UserHandler *UserHandler
	FieldHandler *FieldHandler
}

// func NewServer(handler *UserHandler) *Server {
// 	return &Server{
// 		UserHandler: handler,
// 	}
// }
func NewServer(handler *FieldHandler) *Server {
	return &Server{
		FieldHandler: handler,
	}
}

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/field", s.FieldHandler.ServeHTTP)
	// hacer lo mismo para field

	fmt.Println("listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Failed to start server: %s \n", err)
	}
}
