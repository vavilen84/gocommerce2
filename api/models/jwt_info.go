package models

import (
	"api/constants"
	"api/env"
	"api/helpers"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type JWTInfo struct {
	BaseModel
	User      *User  `orm:"rel(fk)"`
	Secret    string `json:"string"`
	CreatedAt int    `json:"created_at" orm:"column(created_at)"`
	ExpiresAt int    `json:"expires_at" orm:"column(expires_at)"`
}

func (m *JWTInfo) TableName() string {
	return constants.JWTInfoTableName
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

func (m *JWTInfo) generateSecret() {
	m.Secret = helpers.GenerateRandomString(64)
}

func (m *JWTInfo) setTimestampsOnInsert() {
	expiresAt := time.Now().Add(time.Duration(env.GetJWTExpireDurationDays()) * 24 * time.Hour)
	m.CreatedAt = helpers.GetNowUTCTimestamp()
	m.ExpiresAt = int(expiresAt.UTC().Unix())
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
