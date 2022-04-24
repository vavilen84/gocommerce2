package models

const (
	user1key    = "user_1"
	Product1Sku = "product_1_sku"
)

var (
	usersFixtures    map[string]User
	productsFixtures map[int]Product
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
	productsFixtures = map[int]Product{
		1: {
			Title: "Product #1 title",
			SKU:   Product1Sku,
			Price: 100,
		},
	}
}
