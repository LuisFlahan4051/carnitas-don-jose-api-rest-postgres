package models

type Branch struct {
	ID

	Name    string `gorm:"type:varchar(50)" json:"name"`
	Address string `json:"address"`

	BranchSafeboxes     []BranchSafebox      `json:"branch_safeboxes"`
	BranchProductsStock []BranchProductStock `json:"branch_products_stock"`
	BranchSuppliesStock []BranchSupplieStock `json:"branch_suplies_stock"`
	BranchArticlesStock []BranchArticleStock `json:"branch_articles_stock"`
	Users               []User               `json:"users"`
	BranchUserRoles     []BranchUserRole     `json:"branch_user_roles"`
	Turns               []Turn               `json:"turns"`
	Sales               []Sale               `json:"sales"`
	Inventories         []Inventory          `json:"inventories"`
	AdminNotifications  []AdminNotification  `json:"admin_notifications"`
}

type BranchSafebox struct {
	ID

	Name    string  `gorm:"type:varchar(50)" json:"name"`
	Content float64 `gorm:"check: content >= 0; default: 0" json:"content"`

	BranchID  uint `gorm:"not null" json:"branch_id"`
	SafeboxID uint `gorm:"not null; unique" json:"safebox_id"`
}

type BranchProductStock struct {
	ID

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	BranchID  uint `json:"branch_id"`
	ProductID uint `json:"product_id"`
}

type BranchSupplieStock struct {
	ID

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	BranchID uint `json:"branch_id"`
	SupplyID uint `json:"supply_id"`
}

type BranchArticleStock struct {
	ID

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	BranchID  uint `json:"branch_id"`
	ArticleID uint `json:"article_id"`
}
