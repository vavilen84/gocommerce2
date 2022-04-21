package models

import (
	"github.com/anaskhan96/go-password-encoder"
	"github.com/astaxie/beego/orm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_InsertUser(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// validation check
	u := User{}
	err := InsertUser(o, &u)
	assert.Nil(t, err)
	assert.Empty(t, u.Id)
	assert.NotEmpty(t, u.ValidationErrors["email"])
	assert.NotEmpty(t, u.ValidationErrors["password"])
	assert.NotEmpty(t, u.ValidationErrors["first_name"])
	assert.NotEmpty(t, u.ValidationErrors["last_name"])
	assert.NotEmpty(t, u.ValidationErrors["role"])

	u = usersFixtures[user1key]
	rawPassword := u.Password

	// successfully inserted
	err = InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Empty(t, u.ValidationErrors)

	// find inserted user
	userFromDb, err := FindUserById(o, u.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, userFromDb.Id)
	assert.Equal(t, u.Id, userFromDb.Id)
	assert.Equal(t, u.Email, userFromDb.Email)
	assert.NotEmpty(t, userFromDb.Password)
	assert.NotEmpty(t, userFromDb.Salt)
	assert.True(t, password.Verify(rawPassword, userFromDb.Salt, userFromDb.Password, nil))
	assert.Equal(t, u.Role, userFromDb.Role)
	assert.Equal(t, u.FirstName, userFromDb.FirstName)
	assert.Equal(t, u.LastName, userFromDb.LastName)
	assert.NotEmpty(t, userFromDb.CreatedAt)
	assert.NotEmpty(t, userFromDb.UpdatedAt)
	assert.Empty(t, userFromDb.DeletedAt)

	// insert user with the same email
	userWithTheSameEmail := User{Email: u.Email}
	err = InsertUser(o, &userWithTheSameEmail)
	assert.Nil(t, err)
	assert.Empty(t, userWithTheSameEmail.Id)
	assert.NotEmpty(t, userWithTheSameEmail.ValidationErrors["email"])
}

func Test_FindUserById(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	// successfully inserted
	u := usersFixtures[user1key]
	err := InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Empty(t, u.ValidationErrors)

	// user successfully found
	userFromDb, err := FindUserById(o, u.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, userFromDb.Id)
	assert.Equal(t, userFromDb.Id, u.Id)
	assert.Equal(t, userFromDb.Email, u.Email)

	// not existing id. user not found
	notExistingId := int64(999)
	userFromDb, err = FindUserById(o, notExistingId)
	assert.NotNil(t, err)
	assert.Empty(t, userFromDb.Id)
}

func Test_UpdateUser(t *testing.T) {
	beforeEachTest()
	o := orm.NewOrm()

	u := usersFixtures[user1key]
	rawPassword := u.Password

	// successfully inserted
	err := InsertUser(o, &u)
	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Empty(t, u.ValidationErrors)

	// find inserted user
	userFromDb, err := FindUserById(o, u.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, userFromDb.Id)
	assert.Equal(t, u.Id, userFromDb.Id)
	assert.Equal(t, u.Email, userFromDb.Email)
	assert.NotEmpty(t, userFromDb.Password)
	assert.NotEmpty(t, userFromDb.Salt)
	assert.True(t, password.Verify(rawPassword, userFromDb.Salt, userFromDb.Password, nil))
	assert.Equal(t, u.Role, userFromDb.Role)
	assert.Equal(t, u.FirstName, userFromDb.FirstName)
	assert.Equal(t, u.LastName, userFromDb.LastName)
	assert.NotEmpty(t, userFromDb.CreatedAt)
	assert.NotEmpty(t, userFromDb.UpdatedAt)
	assert.Empty(t, userFromDb.DeletedAt)

	// check validation
	emptyUser := User{}
	err = UpdateUser(o, &emptyUser)
	assert.Nil(t, err)
	assert.NotEmpty(t, emptyUser.ValidationErrors["id"])
	assert.NotEmpty(t, emptyUser.ValidationErrors["User"])
	assert.NotEmpty(t, emptyUser.ValidationErrors["email"])
	assert.NotEmpty(t, emptyUser.ValidationErrors["salt"])
	assert.NotEmpty(t, emptyUser.ValidationErrors["password"])
	assert.NotEmpty(t, emptyUser.ValidationErrors["first_name"])
	assert.NotEmpty(t, emptyUser.ValidationErrors["last_name"])
	assert.NotEmpty(t, emptyUser.ValidationErrors["role"])
	assert.NotEmpty(t, emptyUser.ValidationErrors["created_at"])

	// update user
	newEmail := "updated_email@example.com"
	newPassword := "999999"
	newRole := UserRoleAdmin
	newFirstName := "new_first_name"
	newLastName := "new_last_name"

	u.Email = newEmail
	u.Password = newPassword
	u.Role = newRole
	u.FirstName = newFirstName
	u.LastName = newLastName
	err = UpdateUser(o, &u)
	assert.Nil(t, err)
	assert.Empty(t, u.ValidationErrors)

	// find updated user
	userFromDb, err = FindUserById(o, u.Id)
	assert.Nil(t, err)
	assert.NotEmpty(t, userFromDb.Id)
	assert.Equal(t, u.Id, userFromDb.Id)
	assert.Equal(t, newEmail, userFromDb.Email)
	assert.NotEmpty(t, userFromDb.Password)
	assert.NotEmpty(t, userFromDb.Salt)
	assert.True(t, password.Verify(newPassword, userFromDb.Salt, userFromDb.Password, nil))
	assert.Equal(t, newRole, userFromDb.Role)
	assert.Equal(t, newFirstName, userFromDb.FirstName)
	assert.Equal(t, newLastName, userFromDb.LastName)
	assert.NotEmpty(t, userFromDb.CreatedAt)
	assert.Equal(t, u.CreatedAt, userFromDb.CreatedAt)
	assert.NotEmpty(t, userFromDb.UpdatedAt)
	assert.Empty(t, userFromDb.DeletedAt)
}
