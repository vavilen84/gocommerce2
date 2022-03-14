package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createUser(t *testing.T, o orm.Ormer) *User {
	u := usersFixtures[user1key]
	err := InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	return &u
}

func TestInsertUser(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// validation check
	u := User{}
	err := InsertUser(o, &u)
	assert.Nil(t, err)
	assert.Empty(t, u.Id)
	assert.NotEmpty(t, u.ValidationErrors["email"])
	assert.NotEmpty(t, u.ValidationErrors["password"])
	assert.NotEmpty(t, u.ValidationErrors["salt"])
	assert.NotEmpty(t, u.ValidationErrors["first_name"])
	assert.NotEmpty(t, u.ValidationErrors["last_name"])
	assert.NotEmpty(t, u.ValidationErrors["role"])

	// successfully inserted
	createUser(t, o)
}

func TestProduct_FindById(t *testing.T) {
	//beforeEachTest()
	//o := orm.NewOrm()
	//p := createProduct(t, o)
	//
	//modelFromDb := Product{Id: p.Id}
	//err := modelFromDb.FindById(o)
	//assert.Nil(t, err)
	//assert.Equal(t, p.Title, modelFromDb.Title)
	//assert.Equal(t, p.SKU, modelFromDb.SKU)
	//assert.Equal(t, p.Price, modelFromDb.Price)
}
