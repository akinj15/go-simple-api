package database

import (
	"testing"

	"github.com/akinj15/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Notebook", 10.0)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Notebook", product.Name)
}
