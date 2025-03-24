package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/llilyshkall/mathematical-search-engine/internal/postgres"
	repo "github.com/llilyshkall/mathematical-search-engine/internal/repository/postgres"
	"github.com/llilyshkall/mathematical-search-engine/internal/service"
)

func main() {
	db, err := postgres.New()
	if err != nil {
		log.Println("tut")
		log.Println(err)
		return
	}
	postgresRepo := repo.NewExpressionRepository(db)

	s := service.NewService(postgresRepo)
	http.HandleFunc("/", s.HomeHandler)
	http.HandleFunc("/search", s.SearchHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
