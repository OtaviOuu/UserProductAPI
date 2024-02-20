package database

import "github.com/crudGolangAPI/internal/entity"

// Contrato
type UserInterfaceDB interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterfaceDB interface {
	Create(product *entity.Product) error
	// paginação
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
