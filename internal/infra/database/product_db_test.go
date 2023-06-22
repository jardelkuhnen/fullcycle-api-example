package database

import (
	"fmt"
	"github.com/jardelkuhnen/api-example/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)

	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprint("Product ", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
		assert.NoError(t, err)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 1")
	assert.Equal(t, products[9].Name, "Product 10")

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, products, 10)
	assert.Equal(t, products[0].Name, "Product 11")
	assert.Equal(t, products[9].Name, "Product 20")

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Equal(t, products, 3)
	assert.Equal(t, products[0].Name, "Product 21")
	assert.Equal(t, products[2].Name, "Product 23")

}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, "Product 1")
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, "Product 2")
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}
