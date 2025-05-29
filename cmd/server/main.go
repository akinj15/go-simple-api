package main

import (
	"net/http"

	"github.com/akinj15/go-api/configs"
	"github.com/akinj15/go-api/internal/entity"
	"github.com/akinj15/go-api/internal/infra/database"
	"github.com/akinj15/go-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productHandler := handlers.NewProductHandler(database.NewProduct(db))

	r := chi.NewRouter()
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products", productHandler.GetAllProducts)
	r.Get("/products/{id}", productHandler.GetProductByID)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	http.ListenAndServe(":8000", r)
}
