package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_InsertDiscount(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// validation check
	m := Discount{}
	err := InsertDiscount(o, &m)
	assert.Nil(t, err)
	assert.Empty(t, m.Id)
	assert.NotEmpty(t, m.ValidationErrors["title"])
	assert.NotEmpty(t, m.ValidationErrors["amount"])
	assert.NotEmpty(t, m.ValidationErrors["percentage"])
	assert.NotEmpty(t, m.ValidationErrors["type"])

	m = discountFixtures[discount1key]

	// successfully inserted
	err = InsertDiscount(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted Discount
	discountFromDb, err := FindDiscountById(o, m.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, discountFromDb.Id)
	assert.Equal(t, m.Id, discountFromDb.Id)
	assert.Equal(t, m.Title, discountFromDb.Title)
	assert.Equal(t, m.Amount, discountFromDb.Amount)
	assert.Equal(t, m.Percentage, discountFromDb.Percentage)
	assert.NotEmpty(t, discountFromDb.CreatedAt)
	assert.NotEmpty(t, discountFromDb.UpdatedAt)

	// insert Discount with the same title
	discountWithTheSameTitle := Discount{Title: strings.ToLower(discountFromDb.Title)}
	err = InsertDiscount(o, &discountWithTheSameTitle)
	assert.Nil(t, err)
	assert.Empty(t, discountWithTheSameTitle.Id)
	assert.NotEmpty(t, discountWithTheSameTitle.ValidationErrors["title"])
}

func Test_FindDiscountByTitle(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := discountFixtures[discount1key]

	// successfully inserted
	err := InsertDiscount(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted Discount
	discountFromDb, err := FindDiscountByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, discountFromDb.Id)
	assert.Equal(t, m.Id, discountFromDb.Id)
	assert.Equal(t, m.Title, discountFromDb.Title)
	assert.NotEmpty(t, discountFromDb.CreatedAt)
	assert.NotEmpty(t, discountFromDb.UpdatedAt)
}

func Test_DiscountUpdate(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := discountFixtures[discount1key]

	// successfully inserted
	err := InsertDiscount(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted Discount
	discountFromDb, err := FindDiscountByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, discountFromDb.Id)
	assert.Equal(t, m.Id, discountFromDb.Id)
	assert.Equal(t, m.Title, discountFromDb.Title)
	assert.NotEmpty(t, discountFromDb.CreatedAt)
	assert.NotEmpty(t, discountFromDb.UpdatedAt)

	newTitle := "new_title"

	updatedDiscount := Discount{
		BaseModel: BaseModel{
			Id:        discountFromDb.Id,
			CreatedAt: discountFromDb.CreatedAt,
			UpdatedAt: discountFromDb.UpdatedAt,
		},
		Title: newTitle,
	}
	err = UpdateDiscount(o, &updatedDiscount)
	assert.Nil(t, err)
	assert.NotEmpty(t, discountFromDb.CreatedAt)
	assert.NotEmpty(t, discountFromDb.UpdatedAt)
	assert.Equal(t, discountFromDb.CreatedAt, updatedDiscount.CreatedAt)
	assert.Equal(t, newTitle, updatedDiscount.Title)
}

func TestDiscount_Delete(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := discountFixtures[discount1key]

	// successfully inserted
	err := InsertDiscount(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted Discount
	discountFromDb, err := FindDiscountByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, discountFromDb.Id)
	assert.Equal(t, m.Id, discountFromDb.Id)
	assert.Equal(t, m.Title, discountFromDb.Title)
	assert.NotEmpty(t, discountFromDb.CreatedAt)
	assert.NotEmpty(t, discountFromDb.UpdatedAt)

	// remove Discount
	err = DeleteDiscount(o, &discountFromDb)
	assert.Nil(t, err)

	// try to find deleted Discount
	discountFromDb, err = FindDiscountById(o, m.Id)
	assert.NotNil(t, err)
	assert.Empty(t, discountFromDb.Id)
}
