package models

import (
	"api/constants"
	"api/helpers"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_ = iota
	UserRoleCustomer
	UserRoleAdmin
)

type User struct {
	BaseModel
	Email     string `json:"email" orm:"column(email);unique"`
	Password  string `json:"password" orm:"column(password)"`
	Salt      string `json:"salt" orm:"column(salt)"`
	Role      int    `json:"role" orm:"column(role)"`
	FirstName string `json:"first_name" orm:"column(first_name)"`
	LastName  string `json:"last_name" orm:"column(last_name)"`
	CreatedAt int    `json:"created_at" orm:"column(created_at)"`
	UpdatedAt int    `json:"updated_at" orm:"column(updated_at)"`
}

func FindUserByEmail(o orm.Ormer, email string) (*User, error) {
	u := User{}
	err := o.QueryTable(constants.UserModel).Filter("email", email).One(&u)
	if err != nil {
		logs.Error(err)
	}
	return &u, err
}

func FindUserById(o orm.Ormer, id int64) (m User, err error) {
	err = o.QueryTable(constants.UserModel).Filter("id", id).One(&m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func InsertUser(o orm.Ormer, m *User) (err error) {
	m.clearValidationErrors()
	m.validateRawPassword()
	m.encodePassword()
	m.setTimestampsOnInsert()
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

func UpdateUser(o orm.Ormer, m *User) (err error) {
	m.clearValidationErrors()
	m.setPasswordOnUpdate(o)
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

func DeleteUser(o orm.Ormer, m *User) (err error) {
	_, err = FindUserById(o, m.Id)
	if err != nil {
		logs.Error(err)
		return
	}
	_, err = o.Delete(m)
	if err != nil {
		logs.Error(err)
	}
	return
}

func (m *User) setPasswordOnUpdate(o orm.Ormer) {
	oldUser, err := FindUserById(o, m.Id)
	if err != nil {
		logs.Error(err)
	}
	if m.Password != "" {
		m.encodePassword()
	} else {
		m.Password = oldUser.Password
		m.Salt = oldUser.Salt
	}
}

func (m *User) setTimestampsOnInsert() {
	now := helpers.GetNowUTCTimestamp()
	m.CreatedAt = now
	m.UpdatedAt = now
}

func (m *User) setTimestampsOnUpdate() {
	now := helpers.GetNowUTCTimestamp()
	m.UpdatedAt = now
}

func (m *User) encodePassword() {
	salt, encodedPwd := password.Encode(m.Password, nil)
	m.Password = encodedPwd
	m.Salt = salt
}

func (m *User) validateEmailAlreadyInUse(o orm.Ormer, valid *validation.Validation) {
	userFromDb, err := FindUserByEmail(o, m.Email)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
		}
	} else {
		if (userFromDb.Id != 0) && (m.Id != userFromDb.Id) {
			err := valid.SetError("email", "Email is already in use")
			if err != nil {
				logs.Error(err)
			}
		}
	}
}

func (m *User) ValidateUserExists(o orm.Ormer, valid *validation.Validation) {
	_, err := FindUserById(o, m.Id)
	if err != nil {
		if err != orm.ErrNoRows {
			logs.Error(err)
			return
		} else {
			errMsg := fmt.Sprintf("User with id #%d does not exist", m.Id)
			valid.SetError("User", errMsg)
		}
	}
}

func (m *User) validateRawPassword() {
	valid := validation.Validation{}
	valid.Required(m.Password, "password")
	valid.MaxSize(m.Password, 16, "password")
	if valid.HasErrors() {
		m.setValidationErrors(valid.Errors)
		m.logValidationErrors(valid.Errors, constants.UserModel)
	}
}

func (m *User) validateCommonFields(valid *validation.Validation) {
	valid.Required(m.Email, "email")
	valid.MaxSize(m.Email, 255, "email")
	valid.Email(m.Email, "email")

	valid.Required(m.Salt, "salt")
	valid.Required(m.Password, "password")

	valid.Required(m.Role, "role")
	valid.Range(m.Role, UserRoleCustomer, UserRoleAdmin, "role")

	valid.Required(m.FirstName, "first_name")
	valid.MaxSize(m.FirstName, 255, "first_name")

	valid.Required(m.LastName, "last_name")
	valid.MaxSize(m.LastName, 255, "last_name")

	valid.Required(m.CreatedAt, "created_at")
	valid.Required(m.UpdatedAt, "updated_at")
}

func (m *User) validateOnUpdate(o orm.Ormer) bool {
	valid := validation.Validation{}
	valid.Required(m.Id, "id")
	m.ValidateUserExists(o, &valid)
	m.validateCommonFields(&valid)
	m.validateEmailAlreadyInUse(o, &valid)
	if valid.HasErrors() {
		m.setValidationErrors(valid.Errors)
		m.logValidationErrors(valid.Errors, constants.UserModel)
		return false
	}
	return true
}

func (m *User) validateOnInsert(o orm.Ormer) bool {
	valid := validation.Validation{}
	m.validateCommonFields(&valid)
	m.validateEmailAlreadyInUse(o, &valid)
	if valid.HasErrors() {
		m.setValidationErrors(valid.Errors)
		m.logValidationErrors(valid.Errors, constants.UserModel)
		return false
	}
	return true
}
