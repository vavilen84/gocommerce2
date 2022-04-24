package models

import (
	"api/constants"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/core/validation"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Category struct {
	BaseModel
	Title string `json:"title" orm:"column(title)"`
}

func (m *Category) TableName() string {
	return constants.CategoryDBTable
}

func (m *Category) validateOnInsert(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.ValidateTitleAlreadyUsed(o, &valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.CategoryModel)
		return false
	}
	return true
}

func (m *Category) validateCommonFields(valid *validation.Validation) {
	valid.MaxSize(m.Title, 255, "title")
	valid.Required(m.Title, "title")

	valid.Required(m.CreatedAt, "created_at")
	valid.Required(m.UpdatedAt, "updated_at")
}

func (m *Category) validateOnUpdate(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.ValidateCategoryExists(o, &valid)
	m.ValidateTitleAlreadyUsed(o, &valid)
	if valid.HasErrors() {
		m.handleValidationErrors(valid.Errors, constants.CategoryModel)
		return false
	}
	return true
}

func (m *Category) ValidateTitleAlreadyUsed(o orm.Ormer, valid *validation.Validation) {
	categoryFromDb, err := FindCategoryByTitle(o, m.Title)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
		}
	} else {
		if (categoryFromDb.Id != 0) && (m.Id != categoryFromDb.Id) {
			err := valid.SetError("title", "Title is already in use")
			if err != nil {
				logs.Error(err)
			}
		}
	}
}

func (m *Category) setTimestampsOnCreate() {
	now := int(time.Now().Unix())
	m.CreatedAt = now
	m.UpdatedAt = now
}

func InsertCategory(o orm.Ormer, m *Category) (err error) {
	m.clearValidationErrors()
	m.setTimestampsOnCreate()
	isValid := m.validateOnInsert(o)
	if !isValid {
		return
	}
	_, err = o.Insert(m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func FindCategoryByTitle(o orm.Ormer, title string) (m Category, err error) {
	err = o.QueryTable(constants.CategoryDBTable).Filter("title", title).One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func FindCategoryById(o orm.Ormer, id int64) (m Category, err error) {
	err = o.QueryTable(constants.CategoryDBTable).Filter("id", id).One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (m *Category) ValidateCategoryExists(o orm.Ormer, valid *validation.Validation) {
	_, err := FindCategoryById(o, m.Id)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
			return
		} else {
			errMsg := fmt.Sprintf("Category with id #%d does not exist", m.Id)
			valid.SetError("Category", errMsg)
		}
	}
}

func UpdateCategory(o orm.Ormer, m *Category) (err error) {
	m.clearValidationErrors()
	m.setTimestampsOnUpdate()
	isValid := m.validateOnUpdate(o)
	if !isValid {
		return
	}
	_, err = o.Update(m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func DeleteCategory(o orm.Ormer, m *Category) error {
	_, err := FindCategoryById(o, m.Id)
	if err != nil {
		logs.Error(err)
		return err
	}
	_, err = o.Delete(m)
	if err != nil {
		logs.Error(err)
	}
	return nil
}
