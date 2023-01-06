package models

import (
	"time"
)

type ID struct {
	Id        uint       `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type FoodType struct {
	ID

	Name string `gorm:"type:varchar(50)" json:"name"`

	Foods []Food `json:"foods"`
}

type FoodMeat struct {
	ID

	Name  string `gorm:"type:varchar(50)" json:"name"`
	Foods []Food `json:"foods"`
}

type Food struct {
	ID

	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	FoodTypeID uint `json:"food_type_id"`
	FoodMeatID uint `json:"food_meat_id"`

	ProductFoods []ProductFood `json:"product_foods"`
}

type DrinkSize struct {
	ID

	Name string `gorm:"type:varchar(50)" json:"name"`

	Drinks []Drink `json:"drinks"`
}

type DrinkFlavor struct {
	ID

	Name   string  `gorm:"type:varchar(50)" json:"name"`
	Drinks []Drink `json:"drinks"`
}

type Drink struct {
	ID

	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	DrinkSizeID   uint `json:"drink_size_id"`
	DrinkFlavorID uint `json:"drink_flavor_id"`

	ProductDrinks []ProductDrink `json:"product_drinks"`
}

type Product struct {
	ID
	Name        string  `gorm:"type:varchar(50)" json:"name"`
	Description string  `json:"description"`
	Price       float64 `gorm:"check: price >= 0" json:"price"`
	Photo       string  `json:"photo"`

	ProductFoods           []ProductFood           `json:"product_foods"`
	ProductDrinks          []ProductDrink          `json:"product_drinks"`
	BranchProductsStock    []BranchProductStock    `json:"branch_products_stock"`
	SalesProducts          []SaleProduct           `json:"sales_products"`
	InventoryProductsStock []InventoryProductStock `json:"inventory_products_stock"`
}

type ProductFood struct {
	ID

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`

	FoodID    uint `json:"food_id"`
	ProductID uint `json:"product_id"`
}

type ProductDrink struct {
	ID

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`

	DrinkID   uint `json:"drink_id"`
	ProductID uint `json:"product_id"`
}
