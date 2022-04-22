package models

import (
	"api/constants"
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

func InsertJWTInfo(o orm.Ormer, m *JWTInfo) (err error) {
	m.clearValidationErrors()
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
