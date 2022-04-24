package models

const (
	user1key     = "user_1"
	product1key  = "product_1"
	product1Sku  = "product_1_sku"
	category1key = "category_1"
)

var (
	usersFixtures    map[string]User
	productsFixtures map[string]Product
	categoryFixtures map[string]Category
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
	productsFixtures = map[string]Product{
		product1key: {
			Title: "Product #1 title",
			SKU:   product1Sku,
			Price: 100,
		},
	}
	categoryFixtures = map[string]Category{
		category1key: {
			Title: "Category #1 title",
		},
	}
}
