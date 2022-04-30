package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_InsertTax(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// validation check
	m := Tax{}
	err := InsertTax(o, &m)
	assert.Nil(t, err)
	assert.Empty(t, m.Id)
	assert.NotEmpty(t, m.ValidationErrors["title"])
	assert.NotEmpty(t, m.ValidationErrors["amount"])
	assert.NotEmpty(t, m.ValidationErrors["percentage"])

	m = taxFixtures[tax1key]

	// successfully inserted
	err = InsertTax(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted Tax
	taxFromDb, err := FindTaxById(o, m.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, taxFromDb.Id)
	assert.Equal(t, m.Id, taxFromDb.Id)
	assert.Equal(t, m.Title, taxFromDb.Title)
	assert.Equal(t, m.Amount, taxFromDb.Amount)
	assert.Equal(t, m.Percentage, taxFromDb.Percentage)
	assert.NotEmpty(t, taxFromDb.CreatedAt)
	assert.NotEmpty(t, taxFromDb.UpdatedAt)

	// insert Tax with the same title
	TaxWithTheSameTitle := Tax{Title: strings.ToLower(taxFromDb.Title)}
	err = InsertTax(o, &TaxWithTheSameTitle)
	assert.Nil(t, err)
	assert.Empty(t, TaxWithTheSameTitle.Id)
	assert.NotEmpty(t, TaxWithTheSameTitle.ValidationErrors["title"])
}

func Test_FindTaxByTitle(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := taxFixtures[tax1key]

	// successfully inserted
	err := InsertTax(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted Tax
	taxFromDb, err := FindTaxByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, taxFromDb.Id)
	assert.Equal(t, m.Id, taxFromDb.Id)
	assert.Equal(t, m.Title, taxFromDb.Title)
	assert.NotEmpty(t, taxFromDb.CreatedAt)
	assert.NotEmpty(t, taxFromDb.UpdatedAt)
}

func Test_TaxUpdate(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := taxFixtures[tax1key]

	// successfully inserted
	err := InsertTax(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted Tax
	taxFromDb, err := FindTaxByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, taxFromDb.Id)
	assert.Equal(t, m.Id, taxFromDb.Id)
	assert.Equal(t, m.Title, taxFromDb.Title)
	assert.NotEmpty(t, taxFromDb.CreatedAt)
	assert.NotEmpty(t, taxFromDb.UpdatedAt)

	newTitle := "new_title"

	updatedTax := Tax{
		BaseModel: BaseModel{
			Id:        taxFromDb.Id,
			CreatedAt: taxFromDb.CreatedAt,
			UpdatedAt: taxFromDb.UpdatedAt,
		},
		Title: newTitle,
	}
	err = UpdateTax(o, &updatedTax)
	assert.Nil(t, err)
	assert.NotEmpty(t, taxFromDb.CreatedAt)
	assert.NotEmpty(t, taxFromDb.UpdatedAt)
	assert.Equal(t, taxFromDb.CreatedAt, updatedTax.CreatedAt)
	assert.Equal(t, newTitle, updatedTax.Title)
}

func TestTax_Delete(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := taxFixtures[tax1key]

	// successfully inserted
	err := InsertTax(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted Tax
	taxFromDb, err := FindTaxByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, taxFromDb.Id)
	assert.Equal(t, m.Id, taxFromDb.Id)
	assert.Equal(t, m.Title, taxFromDb.Title)
	assert.NotEmpty(t, taxFromDb.CreatedAt)
	assert.NotEmpty(t, taxFromDb.UpdatedAt)

	// remove Tax
	err = DeleteTax(o, &taxFromDb)
	assert.Nil(t, err)

	// try to find deleted Tax
	taxFromDb, err = FindTaxById(o, m.Id)
	assert.NotNil(t, err)
	assert.Empty(t, taxFromDb.Id)
}
