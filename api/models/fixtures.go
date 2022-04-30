package models

import "api/constants"

const (
	user1key = "user_1"

	product1key = "product_1"
	product1Sku = "product_1_sku"

	category1key = "category_1"

	discount1key = "discount_1"
	discount2key = "discount_2"
	discount3key = "discount_3"
)

var (
	usersFixtures    map[string]User
	productsFixtures map[string]Product
	categoryFixtures map[string]Category
	discountFixtures map[string]Discount
)

func initFixtures() {
	discountFixtures = map[string]Discount{
		discount1key: {
			Title:  "cart_discount",
			Amount: 1,
			Type:   constants.DiscountCartType,
		},
		discount2key: {
			Title:      "category_discount",
			Percentage: 10,
			Type:       constants.DiscountCategoryType,
		},
		discount3key: {
			Title:      "product_discount",
			Percentage: 100,
			Type:       constants.DiscountProductType,
		},
	}

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
