package models

import (
	"api/constants"
	"api/helpers"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

type JWTInfo struct {
	BaseModel
	User      *User  `orm:"rel(fk)"`
	Secret    string `json:"string"`
	CreatedAt int    `json:"created_at" orm:"column(created_at)"`
	ExpiresAt int    `json:"expires_at" orm:"column(expires_at)"`
}

func (m *JWTInfo) TableName() string {
	return constants.JWTInfoDBTable
}

func InsertJWTInfo(o orm.Ormer, m *JWTInfo) (err error) {
	m.clearValidationErrors()
	m.setTimestampsOnInsert()
	m.generateSecret()
	isValid := m.validateOnInsert(o)
	if !isValid {
		return
	}
	id, err := o.Insert(m)
	if err != nil {
		logs.Error(err)
		return
	}
	m.Id = id
	return
}

func FindJWTInfoById(o orm.Ormer, id int64) (jwtInfo JWTInfo, err error) {
	err = o.QueryTable(constants.JWTInfoDBTable).Filter("id", id).One(&jwtInfo)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (m *JWTInfo) generateSecret() {
	m.Secret = helpers.GenerateRandomString(64)
}

func (m *JWTInfo) setTimestampsOnInsert() {
	m.CreatedAt = helpers.GetNowUTCTimestamp()
}

func (m *JWTInfo) clearValidationErrors() {
	m.ValidationErrors = make(map[string][]string)
}

func (m *JWTInfo) validateOnInsert(o orm.Ormer) bool {
	valid := validation.Validation{}

	valid.Required(m.User, "user")
	valid.Required(m.Secret, "secret")
	valid.Required(m.CreatedAt, "created_at")
	valid.Required(m.ExpiresAt, "expires_at")

	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.JWTInfoModel)
		return false
	}

	m.User.ValidateUserExists(o, &valid)

	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.JWTInfoModel)
		return false
	}
	return true
}
