package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddCategory2Product(t *testing.T) {
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

	// validation check
	m := Category{}
	err = InsertCategory(o, &m)
	assert.Nil(t, err)
	assert.Empty(t, m.Id)
	assert.NotEmpty(t, m.ValidationErrors["title"])

	m = categoryFixtures[category1key]

	// successfully inserted
	err = InsertCategory(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted category
	categoryFromDb, err := FindCategoryById(o, m.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, categoryFromDb.Id)
	assert.Equal(t, m.Id, categoryFromDb.Id)
	assert.Equal(t, m.Title, categoryFromDb.Title)
	assert.NotEmpty(t, categoryFromDb.CreatedAt)
	assert.NotEmpty(t, categoryFromDb.UpdatedAt)

	// add category to product
	m2m := o.QueryM2M(&productFromDb, "Categories")
	_, err = m2m.Add(&categoryFromDb)
	assert.Nil(t, err)

	// get related product categories
	productFromDb, err = FindProductById(o, p.Id)
	assert.NotEmpty(t, productFromDb.Categories)
	assert.Equal(t, len(productFromDb.Categories), 1)
	assert.Equal(t, categoryFromDb.Id, productFromDb.Categories[0].Id)
	assert.Equal(t, categoryFromDb.Title, productFromDb.Categories[0].Title)

	// get related category products
	categoryFromDb, err = FindCategoryById(o, m.Id)
	assert.NotEmpty(t, categoryFromDb.Products)
	assert.Equal(t, len(categoryFromDb.Products), 1)
	assert.Equal(t, productFromDb.Id, categoryFromDb.Products[0].Id)
	assert.Equal(t, productFromDb.Title, categoryFromDb.Products[0].Title)
}
