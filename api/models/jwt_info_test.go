package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InsertJWTInfo(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	jwtInfo := JWTInfo{}
	err := InsertJWTInfo(o, &jwtInfo)
	assert.Nil(t, err)
	assert.Empty(t, jwtInfo.Id)
	assert.NotEmpty(t, jwtInfo.ValidationErrors["user"])

	// create user
	u := usersFixtures[user1key]
	err = InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Empty(t, u.ValidationErrors)

	jwtInfo.User = &u
	err = InsertJWTInfo(o, &jwtInfo)
	assert.Nil(t, err)
	assert.NotEmpty(t, jwtInfo.Id)
	assert.Empty(t, jwtInfo.ValidationErrors)
}
