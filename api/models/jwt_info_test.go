package models

import (
	"api/helpers"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
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
	assert.NotEmpty(t, jwtInfo.ValidationErrors["expires_at"])

	// create user
	u := usersFixtures[user1key]
	err = InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Empty(t, u.ValidationErrors)

	jwtInfo.User = &u
	jwtInfo.ExpiresAt = helpers.GetDefaultJWTExpiresAt()
	err = InsertJWTInfo(o, &jwtInfo)
	assert.Nil(t, err)
	assert.NotEmpty(t, jwtInfo.Id)
	assert.Empty(t, jwtInfo.ValidationErrors)

	// find created jwt info
	m, err := FindJWTInfoById(o, jwtInfo.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.Equal(t, jwtInfo.Id, m.Id)
	assert.Equal(t, jwtInfo.User.Id, m.User.Id)
	assert.Equal(t, jwtInfo.Secret, m.Secret)
	assert.Equal(t, jwtInfo.CreatedAt, m.CreatedAt)
	assert.Equal(t, jwtInfo.ExpiresAt, m.ExpiresAt)
}
