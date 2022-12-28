package models

type Supply struct {
	ID

	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	BranchSuppliesStock    []BranchSupplieStock    `json:"branch_suplies_stock"`
	InventorySuppliesStock []InventorySupplieStock `json:"inventory_supplies_stock"`
}

type Article struct {
	ID

	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	BranchArticlesStock    []BranchArticleStock    `json:"branch_articles_stock"`
	InventoryArticlesStock []InventoryArticleStock `json:"inventory_articles_stock"`
}

type Safebox struct {
	ID

	Cents10 uint `gorm:"check: cents10 >= 0; default: 0" json:"cents10"`
	Cents50 uint `gorm:"check: cents50 >= 0; default: 0" json:"cents50"`
	Coins1  uint `gorm:"check: coins1 >= 0; default: 0" json:"coins1"`
	Coins2  uint `gorm:"check: coins2 >= 0; default: 0" json:"coins2"`
	Coins5  uint `gorm:"check: coins5 >= 0; default: 0" json:"coins5"`
	Coins10 uint `gorm:"check: coins10 >= 0; default: 0" json:"coins10"`
	Coins20 uint `gorm:"check: coins20 >= 0; default: 0" json:"coins20"`

	Bills20   uint `gorm:"check: bills20 >= 0; default: 0" json:"bills20"`
	Bills50   uint `gorm:"check: bills50 >= 0; default: 0" json:"bills50"`
	Bills100  uint `gorm:"check: bills100 >= 0; default: 0" json:"bills100"`
	Bills200  uint `gorm:"check: bills200 >= 0; default: 0" json:"bills200"`
	Bills500  uint `gorm:"check: bills500 >= 0; default: 0" json:"bills500"`
	Bills1000 uint `gorm:"check: bills1000 >= 0; default: 0" json:"bills1000"`

	BranchSafeboxes    []BranchSafebox    `json:"branch_safeboxes"`
	TurnSafeboxes      []TurnSafebox      `json:"turn_safeboxes"`
	InventorySafeboxes []InventorySafebox `json:"inventory_safeboxes"`
	ActionSafebox      `json:"action_safebox"`
}

type Income struct {
	ID

	Reason      string       `json:"reason"`
	Income      float64      `gorm:"check: income >= 0; not null" json:"income"`
	SaleIncomes []SaleIncome `json:"sale_incomes"`
}

type Expense struct {
	ID

	Reason       string        `json:"reason"`
	Expense      float64       `gorm:"check: expense >= 0; not null" json:"expense"`
	SaleExpenses []SaleExpense `json:"sale_expenses"`
}

type Argument struct {
	ID

	Complaint     bool           `json:"complaint"`
	Score         uint           `gorm:"check: score >= 0; default: 0" json:"score"`
	Argument      string         `json:"argument"`
	SaleArguments []SaleArgument `json:"sale_arguments"`
}
