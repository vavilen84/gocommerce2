package services

import "api/models"

const (
	user1key = "user_1"
)

var (
	usersFixtures map[string]models.User
)

func initFixtures() {
	usersFixtures = map[string]models.User{
		user1key: {
			Email:     "user_1@example.com",
			Password:  "123456",
			Role:      models.UserRoleCustomer,
			FirstName: "John",
			LastName:  "Dou",
		},
	}
}
