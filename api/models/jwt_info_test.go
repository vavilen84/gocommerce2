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
	assert.NotEmpty(t, jwtInfo.ValidationErrors["secret"])
	assert.NotEmpty(t, jwtInfo.ValidationErrors["created_at"])
	assert.NotEmpty(t, jwtInfo.ValidationErrors["expires_at"])
}
