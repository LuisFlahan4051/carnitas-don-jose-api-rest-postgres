package models

type InventoryType struct {
	ID

	Type        string      `json:"type"`
	Inventories []Inventory `json:"inventories"`
}

type Inventory struct {
	ID

	Acepted bool `json:"acepted"`

	InventoryTypeID uint `json:"inventory_type_id"` // INIT/FINAL/LOSSES
	BranchID        uint `json:"branch_id"`
	TurnID          uint `json:"turn_id"`

	InventoryProductsStocks []InventoryProductStock `json:"inventory_products_stocks"`
	InventorySuppliesStocks []InventorySupplyStock  `json:"inventory_supplies_stocks"`
	InventoryArticlesStocks []InventoryArticleStock `json:"inventory_articles_stocks"`
	InventorySafebox        `json:"inventory_safebox"`
}

type InventoryProductStock struct {
	ID

	UnitQuantity     uint `json:"unit_quantity"`
	GrammageQuantity uint `json:"grammage_quantity"`
	InUse            bool `json:"inUse"`

	InventoryID uint `json:"inventory_id"`
	ProductID   uint `json:"product_id"`
}

type InventorySupplyStock struct {
	ID

	UnitQuantity     uint `json:"unit_quantity"`
	GrammageQuantity uint `json:"grammage_quantity"`
	InUse            bool `json:"inUse"`

	InventoryID uint `json:"inventory_id"`
	SupplyID    uint `json:"supply_id"`
}

type InventoryArticleStock struct {
	ID

	UnitQuantity     uint `json:"unit_quantity"`
	GrammageQuantity uint `json:"grammage_quantity"`
	InUse            bool `json:"inUse"`

	InventoryID uint `json:"inventory_id"`
	ArticleID   uint `json:"article_id"`
}

type InventorySafebox struct {
	ID

	InventoryID uint `json:"inventory_id"`
	SafeboxID   uint `json:"safebox_id"`
}
