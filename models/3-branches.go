package models

type Branch struct {
	ID

	Name    string `json:"name"`
	Address string `json:"address"`

	BranchSafeboxes     []BranchSafebox      `json:"safeboxes,omitempty"`
	BranchProductsStock []BranchProductStock `json:"branch_products_stock,omitempty"`
	BranchSuppliesStock []BranchSupplyStock  `json:"branch_suplies_stock,omitempty"`
	BranchArticlesStock []BranchArticleStock `json:"branch_articles_stock,omitempty"`
	Users               []User               `json:"users,omitempty"`
	BranchUserRoles     []BranchUserRole     `json:"branch_user_roles,omitempty"`
	Turns               []Turn               `json:"turns,omitempty"`
	Sales               []Sale               `json:"sales,omitempty"`
	Inventories         []Inventory          `json:"inventories,omitempty"`
	AdminNotifications  []AdminNotification  `json:"admin_notifications,omitempty"`
}

type BranchSafebox struct {
	ID

	Name    string   `json:"name"`
	Content *float64 `json:"content,omitempty"`

	BranchID  uint `json:"branch_id"`
	SafeboxID uint `json:"safebox_id"`
}

type BranchProductStock struct {
	ID

	UnitQuantity     uint  `json:"unit_quantity"`
	GrammageQuantity uint  `json:"grammage_quantity"`
	InUse            *bool `json:"in_use,omitempty"`

	BranchID  uint `json:"branch_id"`
	ProductID uint `json:"product_id"`
}

type BranchSupplyStock struct {
	ID

	UnitQuantity     uint  `json:"unit_quantity"`
	GrammageQuantity uint  `json:"grammage_quantity"`
	InUse            *bool `json:"in_use,omitempty"`

	BranchID uint `json:"branch_id"`
	SupplyID uint `json:"supply_id"`
}

type BranchArticleStock struct {
	ID

	UnitQuantity     uint  `json:"unit_quantity"`
	GrammageQuantity uint  `json:"grammage_quantity"`
	InUse            *bool `json:"in_use,omitempty"`

	BranchID  uint `json:"branch_id"`
	ArticleID uint `json:"article_id"`
}
