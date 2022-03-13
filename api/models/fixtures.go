package models

const (
	user1key = "user_1"
)

var (
	usersFixtures map[string]User
)

func initFixtures() {
	usersFixtures = map[string]User{
		user1key: {
			Email:     "user_1@example.com",
			Password:  "123456",
			Role:      UserRoleCustomer,
			FirstName: "John",
			LastName:  "Dou",
		},
	}
}
