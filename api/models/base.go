package models

import (
	"api/constants"
	"api/helpers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(JWTInfo))
	orm.RegisterModel(new(Product))
	orm.RegisterModel(new(Category))
}

type BaseModel struct {
	Id               int64               `json:"id" orm:"auto"`
	ValidationErrors map[string][]string `orm:"-"`
	CreatedAt        int                 `json:"created_at" orm:"column(created_at)"`
	UpdatedAt        int                 `json:"updated_at" orm:"column(updated_at)"`
}

func (m *BaseModel) clearValidationErrors() {
	m.ValidationErrors = make(map[string][]string)
}

func (m *BaseModel) handleValidationErrors(errors []*validation.Error, modelName string) {
	m.setValidationErrors(errors)
	m.logValidationErrors(errors, constants.JWTInfoModel)
}

func (m *BaseModel) setValidationErrors(errors []*validation.Error) {
	if errors == nil || len(errors) == 0 {
		return
	}
	if len(m.ValidationErrors) == 0 {
		m.ValidationErrors = make(map[string][]string)
	}
	for _, err := range errors {
		if _, ok := m.ValidationErrors[err.Key]; !ok {
			m.ValidationErrors[err.Key] = make([]string, 1)
			m.ValidationErrors[err.Key][0] = err.Message
		} else {
			m.ValidationErrors[err.Key] = append(m.ValidationErrors[err.Key], err.Message)
		}
	}
}

func (m *BaseModel) logValidationErrors(errors []*validation.Error, modelName string) {
	for _, err := range errors {
		logs.Error("Validation error; Model: %v; Key: %v; Message: %v", modelName, err.Key, err.Message)
	}
}

func (m *BaseModel) setTimestampsOnInsert() {
	now := helpers.GetNowUTCTimestamp()
	m.CreatedAt = now
	m.UpdatedAt = now
}

func (m *BaseModel) setTimestampsOnUpdate() {
	now := helpers.GetNowUTCTimestamp()
	m.UpdatedAt = now
}
