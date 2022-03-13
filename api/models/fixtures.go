package models

const (
	Product1Sku = "product_1_sku"
)

var (
	usersFixtures map[int]User
)

func initFixtures() {
	usersFixtures = map[int]User{
		1: {
			Email:     "user_1@example.com",
			Password:  "123456",
			Role:      UserRoleCustomer,
			FirstName: "John",
			LastName:  "Dou",
		},
	}
}
