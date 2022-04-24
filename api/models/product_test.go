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

//func TestProduct_FindById(t *testing.T) {
//	beforeEachTest()
//	o := orm.NewOrm()
//	p := createProduct(t, o)
//
//	modelFromDb := Product{Id: p.Id}
//	err := modelFromDb.FindById(o)
//	assert.Nil(t, err)
//	assert.Equal(t, p.Title, modelFromDb.Title)
//	assert.Equal(t, p.SKU, modelFromDb.SKU)
//	assert.Equal(t, p.Price, modelFromDb.Price)
//}
//
//func TestProduct_FindByDKU(t *testing.T) {
//	beforeEachTest()
//	o := orm.NewOrm()
//	p := createProduct(t, o)
//
//	modelFromDb := Product{SKU: p.SKU}
//	err := modelFromDb.FindBySKU(o)
//	assert.Nil(t, err)
//	assert.Equal(t, p.Title, modelFromDb.Title)
//	assert.Equal(t, p.SKU, modelFromDb.SKU)
//	assert.Equal(t, p.Price, modelFromDb.Price)
//}
//
//func TestProduct_Update(t *testing.T) {
//	beforeEachTest()
//	o := orm.NewOrm()
//	p := createProduct(t, o)
//	newTitle := "New Title"
//	p.Title = newTitle
//	err := p.Update(o)
//	assert.Nil(t, err)
//
//	modelFromDb := Product{Id: p.Id}
//	err = modelFromDb.FindById(o)
//	assert.Nil(t, err)
//	assert.Equal(t, newTitle, modelFromDb.Title)
//}
//
//func TestProduct_Delete(t *testing.T) {
//	beforeEachTest()
//	o := orm.NewOrm()
//	p := createProduct(t, o)
//
//	modelFromDb := Product{Id: p.Id}
//	err := modelFromDb.FindById(o)
//	assert.Nil(t, err)
//
//	err = modelFromDb.Delete(o)
//	assert.Nil(t, err)
//
//	deletedProduct := Product{Id: p.Id}
//	err = deletedProduct.FindById(o)
//	assert.NotNil(t, err)
//	assert.Equal(t, orm.ErrNoRows, err)
//}
