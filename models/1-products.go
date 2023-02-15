package models

import (
	"time"
)

type ID struct {
	Id        uint       `json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type FoodType struct {
	ID

	Name string `json:"name"`

	Foods []Food `json:"foods,omitempty"`
}

type FoodMeat struct {
	ID

	Name  string `json:"name"`
	Foods []Food `json:"foods,omitempty"`
}

type Food struct {
	ID

	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Photo       *string `json:"photo,omitempty"`

	FoodTypeID uint `json:"food_type_id"`
	FoodMeatID uint `json:"food_meat_id"`

	ProductFoods []ProductFood `json:"product_foods,omitempty"`
}

type DrinkSize struct {
	ID

	Name string `json:"name"`

	Drinks []Drink `json:"drinks,omitempty"`
}

type DrinkFlavor struct {
	ID

	Name   string  `json:"name"`
	Drinks []Drink `json:"drinks,omitempty"`
}

type Drink struct {
	ID

	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	DrinkSizeID   uint `json:"drink_size_id"`
	DrinkFlavorID uint `json:"drink_flavor_id"`

	ProductDrinks []ProductDrink `json:"product_drinks,omitempty"`
}

type Product struct {
	ID
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Photo       *string `json:"photo,omitempty"`

	ProductFoods           []ProductFood           `json:"product_foods,omitempty"`
	ProductDrinks          []ProductDrink          `json:"product_drinks,omitempty"`
	BranchProductsStock    []BranchProductStock    `json:"branch_products_stock,omitempty"`
	SalesProducts          []SaleProduct           `json:"sales_products,omitempty"`
	InventoryProductsStock []InventoryProductStock `json:"inventory_products_stock,omitempty"`
}

type ProductFood struct {
	ID

	UnitQuantity     uint `json:"unit_quantity"`
	GrammageQuantity uint `json:"grammage_quantity"`

	FoodID    uint `json:"food_id"`
	ProductID uint `json:"product_id"`
}

type ProductDrink struct {
	ID

	UnitQuantity     uint `json:"unit_quantity"`
	GrammageQuantity uint `json:"grammage_quantity"`

	DrinkID   uint `json:"drink_id"`
	ProductID uint `json:"product_id"`
}
