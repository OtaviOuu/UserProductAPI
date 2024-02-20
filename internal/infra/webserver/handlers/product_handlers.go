package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/crudGolangAPI/internal/dto"
	"github.com/crudGolangAPI/internal/entity"
	"github.com/crudGolangAPI/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterfaceDB
}

func NewProductHandler(db database.ProductInterfaceDB) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
