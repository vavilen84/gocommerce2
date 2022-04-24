package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InsertProduct(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// validation check
	p := Product{}
	err := InsertProduct(o, &p)
	assert.Nil(t, err)
	assert.Empty(t, p.Id)
	assert.NotEmpty(t, p.ValidationErrors["sku"])
	assert.NotEmpty(t, p.ValidationErrors["title"])
	assert.NotEmpty(t, p.ValidationErrors["price"])

	p = productsFixtures[product1key]

	// successfully inserted
	err = InsertProduct(o, &p)
	assert.Nil(t, err)
	assert.NotEmpty(t, p.Id)
	assert.Empty(t, p.ValidationErrors)

	// find inserted product
	productFromDb, err := FindProductById(o, p.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, productFromDb.Id)
	assert.Equal(t, p.Id, productFromDb.Id)
	assert.Equal(t, p.Title, productFromDb.Title)
	assert.Equal(t, p.SKU, productFromDb.SKU)
	assert.NotEmpty(t, productFromDb.CreatedAt)
	assert.NotEmpty(t, productFromDb.UpdatedAt)

	// insert product with the same sku
	productWithTheSameSKU := Product{SKU: productFromDb.SKU}
	err = InsertProduct(o, &productWithTheSameSKU)
	assert.Nil(t, err)
	assert.Empty(t, productWithTheSameSKU.Id)
	assert.NotEmpty(t, productWithTheSameSKU.ValidationErrors["sku"])
}

func Test_FindProductBySKU(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	p := productsFixtures[product1key]

	// successfully inserted
	err := InsertProduct(o, &p)
	assert.Nil(t, err)
	assert.NotEmpty(t, p.Id)
	assert.Empty(t, p.ValidationErrors)

	// find inserted product
	productFromDb, err := FindProductBySKU(o, p.SKU)
	assert.Nil(t, err)
	assert.NotEmpty(t, productFromDb.Id)
	assert.Equal(t, p.Id, productFromDb.Id)
	assert.Equal(t, p.Title, productFromDb.Title)
	assert.Equal(t, p.SKU, productFromDb.SKU)
	assert.NotEmpty(t, productFromDb.CreatedAt)
	assert.NotEmpty(t, productFromDb.UpdatedAt)
}

func Test_ProductUpdate(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	p := productsFixtures[product1key]

	// successfully inserted
	err := InsertProduct(o, &p)
	assert.Nil(t, err)
	assert.NotEmpty(t, p.Id)
	assert.Empty(t, p.ValidationErrors)

	// find inserted product
	productFromDb, err := FindProductById(o, p.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, productFromDb.Id)
	assert.Equal(t, p.Id, productFromDb.Id)
	assert.Equal(t, p.Title, productFromDb.Title)
	assert.Equal(t, p.SKU, productFromDb.SKU)
	assert.NotEmpty(t, productFromDb.CreatedAt)
	assert.NotEmpty(t, productFromDb.UpdatedAt)

	newTitle := "new_title"
	newSKU := "new_sku"
	newPrice := 200

	updatedProduct := Product{
		BaseModel: BaseModel{
			Id:        productFromDb.Id,
			CreatedAt: productFromDb.CreatedAt,
			UpdatedAt: productFromDb.UpdatedAt,
		},
		Title: newTitle,
		SKU:   newSKU,
		Price: newPrice,
	}
	err = UpdateProduct(o, &updatedProduct)
	assert.Nil(t, err)
	assert.NotEmpty(t, productFromDb.CreatedAt)
	assert.NotEmpty(t, productFromDb.UpdatedAt)
	assert.Equal(t, productFromDb.CreatedAt, updatedProduct.CreatedAt)
	assert.Equal(t, newTitle, updatedProduct.Title)
	assert.Equal(t, newSKU, updatedProduct.SKU)
	assert.Equal(t, newPrice, updatedProduct.Price)
}

func TestProduct_Delete(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	p := productsFixtures[product1key]

	// successfully inserted
	err := InsertProduct(o, &p)
	assert.Nil(t, err)
	assert.NotEmpty(t, p.Id)
	assert.Empty(t, p.ValidationErrors)

	// find inserted product
	productFromDb, err := FindProductBySKU(o, p.SKU)
	assert.Nil(t, err)
	assert.NotEmpty(t, productFromDb.Id)
	assert.Equal(t, p.Id, productFromDb.Id)
	assert.Equal(t, p.Title, productFromDb.Title)
	assert.Equal(t, p.SKU, productFromDb.SKU)
	assert.NotEmpty(t, productFromDb.CreatedAt)
	assert.NotEmpty(t, productFromDb.UpdatedAt)

	// remove product
	err = DeleteProduct(o, &productFromDb)
	assert.Nil(t, err)

	// try to find deleted product
	productFromDb, err = FindProductBySKU(o, p.SKU)
	assert.NotNil(t, err)
	assert.Empty(t, productFromDb.Id)
}
