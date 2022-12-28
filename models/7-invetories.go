package models

type InventoryType struct {
	ID

	InventoryType string      `json:"inventory_type"`
	Inventories   []Inventory `json:"inventories"`
}

type Inventory struct {
	ID

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	InventoryTypeID uint `json:"inventory_type_id"` // INIT/FINAL/LOSSES
	BranchID        uint `json:"branch_id"`
	TurnID          uint `json:"turn_id"`

	InventoryProductsStocks []InventoryProductStock `json:"inventory_products_stocks"`
	InventorySuppliesStocks []InventorySupplieStock `json:"inventory_supplies_stocks"`
	InventoryArticlesStocks []InventoryArticleStock `json:"inventory_articles_stocks"`
	InventorySafebox        `json:"inventory_safebox"`
}

type InventoryProductStock struct {
	ID

	InventoryID uint `json:"inventory_id"`
	ProductID   uint `json:"product_id"`
}

type InventorySupplieStock struct {
	ID

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"inUse"`

	InventoryID uint `json:"inventory_id"`
	SupplyID    uint `json:"supply_id"`
}

type InventoryArticleStock struct {
	ID

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	InventoryID uint `json:"inventory_id"`
	ArticleID   uint `json:"article_id"`
}

type InventorySafebox struct {
	ID

	InventoryID uint `gorm:"unique" json:"inventory_id"`
	SafeboxID   uint `gorm:"unique" json:"safebox_id"`
}
