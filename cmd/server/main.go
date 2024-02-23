package main

import (
	"net/http"

	"github.com/crudGolangAPI/internal/entity"
	"github.com/crudGolangAPI/internal/infra/database"
	"github.com/crudGolangAPI/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// _, err := configs.LoadConfig(".env")
	// if err != nil {
	// 	panic(err)
	// }

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.Product{})
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	userDB, _ := database.NewUserDB(db)
	productDB, _ := database.NewProductDB(db)

	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/products/{id}", productHandler.CreateProduct)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	r.Post("/user", userHandler.Create)

	http.ListenAndServe(":8000", r)

}
