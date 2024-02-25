package main

import (
	"net/http"

	"github.com/crudGolangAPI/configs"
	"github.com/crudGolangAPI/internal/entity"
	"github.com/crudGolangAPI/internal/infra/database"
	"github.com/crudGolangAPI/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2" // http-swagger middleware
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	cfg := configs.LoadConfig()

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.Product{})
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	userDB, _ := database.NewUserDB(db)
	productDB, _ := database.NewProductDB(db)

	productHandler := handlers.NewProductHandler(productDB)
	userHandler := handlers.NewUserHandler(userDB, cfg.TokenAuth, cfg.JWTExperesIn)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/products", func(chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth)) // Valida token (olha o secret)
		r.Get("/{id}", productHandler.CreateProduct)
		r.Post("/", productHandler.CreateProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/user/generate_token", userHandler.GetJwt)

	r.Post("/user", userHandler.Create)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost"+cfg.WebServerPort+"/swagger/doc.json"),
	))

	http.ListenAndServe(cfg.WebServerPort, r)

}
