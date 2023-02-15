package models

type Supply struct {
	ID

	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Photo       *string `json:"photo,omitempty"`

	BranchSuppliesStock    []BranchSupplyStock    `json:"branch_suplies_stock,omitempty"`
	InventorySuppliesStock []InventorySupplyStock `json:"inventory_supplies_stock,omitempty"`
}

type Article struct {
	ID

	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Photo       *string `json:"photo,omitempty"`

	BranchArticlesStock    []BranchArticleStock    `json:"branch_articles_stock,omitempty"`
	InventoryArticlesStock []InventoryArticleStock `json:"inventory_articles_stock,omitempty"`
}

type Safebox struct {
	ID

	Cents10 *uint `json:"cents10,omitempty"`
	Cents50 *uint `json:"cents50,omitempty"`
	Coins1  *uint `json:"coins1,omitempty"`
	Coins2  *uint `json:"coins2,omitempty"`
	Coins5  *uint `json:"coins5,omitempty"`
	Coins10 *uint `json:"coins10,omitempty"`
	Coins20 *uint `json:"coins20,omitempty"`

	Bills20   *uint `json:"bills20,omitempty"`
	Bills50   *uint `json:"bills50,omitempty"`
	Bills100  *uint `json:"bills100,omitempty"`
	Bills200  *uint `json:"bills200,omitempty"`
	Bills500  *uint `json:"bills500,omitempty"`
	Bills1000 *uint `json:"bills1000,omitempty"`

	BranchSafeboxes    []BranchSafebox    `json:"branch_safeboxes,omitempty"`
	TurnSafeboxes      []TurnSafebox      `json:"turn_safeboxes,omitempty"`
	InventorySafeboxes []InventorySafebox `json:"inventory_safeboxes,omitempty"`
	ActionSafebox      []ActionSafebox    `json:"action_safebox,omitempty"`
}

type Income struct {
	ID

	Reason      string       `json:"reason"`
	Income      float64      `json:"income"`
	SaleIncomes []SaleIncome `json:"sale_incomes,omitempty"`
}

type Expense struct {
	ID

	Reason       string        `json:"reason"`
	Expense      float64       `json:"expense"`
	SaleExpenses []SaleExpense `json:"sale_expenses,omitempty"`
}

type Argument struct {
	ID

	Complaint     bool           `json:"complaint"`
	Score         uint           `json:"score"`
	Argument      string         `json:"argument"`
	SaleArguments []SaleArgument `json:"sale_arguments,omitempty"`
}
