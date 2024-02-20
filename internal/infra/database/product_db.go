package database

import (
	"github.com/crudGolangAPI/internal/entity"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProductDB(db *gorm.DB) (*ProductDB, error) {
	return &ProductDB{
		DB: db,
	}, nil
}

func (p *ProductDB) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductDB) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("Created_At" + sort).Find(&products).Error
	} else {
		err = p.DB.Order("Created_At" + sort).Find(&products).Error
	}

	return products, err

}

func (p *ProductDB) FindById(id string) (*entity.Product, error) {
	var foundProduct entity.Product
	err := p.DB.Find(&foundProduct, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &foundProduct, nil

}

func (p *ProductDB) Update(product *entity.Product) error {
	// Confirmar se o usuario existe no banco antes de atualizar
	_, err := p.FindById(product.ID.String())
	if err != nil {
		return err
	}

	return p.DB.Save(product).Error

}

func (p *ProductDB) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
