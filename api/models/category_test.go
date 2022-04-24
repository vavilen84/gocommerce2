package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InsertCategory(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// validation check
	m := Category{}
	err := InsertCategory(o, &m)
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

	// insert category with the same title
	categoryWithTheSameTitle := Category{Title: categoryFromDb.Title}
	err = InsertCategory(o, &categoryWithTheSameTitle)
	assert.Nil(t, err)
	assert.Empty(t, categoryWithTheSameTitle.Id)
	assert.NotEmpty(t, categoryWithTheSameTitle.ValidationErrors["title"])
}

func Test_FindCategoryByTitle(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := categoryFixtures[category1key]

	// successfully inserted
	err := InsertCategory(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted category
	categoryFromDb, err := FindCategoryByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, categoryFromDb.Id)
	assert.Equal(t, m.Id, categoryFromDb.Id)
	assert.Equal(t, m.Title, categoryFromDb.Title)
	assert.NotEmpty(t, categoryFromDb.CreatedAt)
	assert.NotEmpty(t, categoryFromDb.UpdatedAt)
}

func Test_CategoryUpdate(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := categoryFixtures[category1key]

	// successfully inserted
	err := InsertCategory(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted category
	categoryFromDb, err := FindCategoryByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, categoryFromDb.Id)
	assert.Equal(t, m.Id, categoryFromDb.Id)
	assert.Equal(t, m.Title, categoryFromDb.Title)
	assert.NotEmpty(t, categoryFromDb.CreatedAt)
	assert.NotEmpty(t, categoryFromDb.UpdatedAt)

	newTitle := "new_title"

	updatedCategory := Category{
		BaseModel: BaseModel{
			Id:        categoryFromDb.Id,
			CreatedAt: categoryFromDb.CreatedAt,
			UpdatedAt: categoryFromDb.UpdatedAt,
		},
		Title: newTitle,
	}
	err = UpdateCategory(o, &updatedCategory)
	assert.Nil(t, err)
	assert.NotEmpty(t, categoryFromDb.CreatedAt)
	assert.NotEmpty(t, categoryFromDb.UpdatedAt)
	assert.Equal(t, categoryFromDb.CreatedAt, updatedCategory.CreatedAt)
	assert.Equal(t, newTitle, updatedCategory.Title)
}

func TestCategory_Delete(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	m := categoryFixtures[category1key]

	// successfully inserted
	err := InsertCategory(o, &m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Empty(t, m.ValidationErrors)

	// find inserted category
	categoryFromDb, err := FindCategoryByTitle(o, m.Title)
	assert.Nil(t, err)
	assert.NotEmpty(t, categoryFromDb.Id)
	assert.Equal(t, m.Id, categoryFromDb.Id)
	assert.Equal(t, m.Title, categoryFromDb.Title)
	assert.NotEmpty(t, categoryFromDb.CreatedAt)
	assert.NotEmpty(t, categoryFromDb.UpdatedAt)

	// remove category
	err = DeleteCategory(o, &categoryFromDb)
	assert.Nil(t, err)

	// try to find deleted category
	categoryFromDb, err = FindCategoryById(o, m.Id)
	assert.NotNil(t, err)
	assert.Empty(t, categoryFromDb.Id)
}
