package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createUser(t *testing.T, o orm.Ormer) *User {
	u := usersFixtures[1]
	err := InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	return &u
}

func TestInsertUser(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()
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
