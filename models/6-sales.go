package models

import (
	"time"
)

type Sale struct {
	ID

	EntryDate   time.Time `json:"entry_date"`
	ExitDate    time.Time `json:"exit_date"`
	TableNumber uint      `json:"table_number"`
	Packed      bool      `json:"packed"`

	UserID         uint           `json:"user_id"`
	BranchID       uint           `json:"branch_id"`
	TurnID         uint           `json:"turn_id"`
	SalesProducts  []SaleProduct  `json:"sales_products"`
	SalesIncomes   []SaleIncome   `json:"sales_incomes"`
	SalesExpenses  []SaleExpense  `json:"sales_expenses"`
	SalesArguments []SaleArgument `json:"sales_arguments"`
}

type SaleProduct struct {
	ID

	Done              bool    `gorm:"default: false" json:"done"`
	GrammageQuantity  uint    `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	KilogrammagePrice float64 `gorm:"check: kilogrammage_price >= 0; default: 0" json:"kilogrammage_price"`
	UnitQuantity      uint    `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	UnitPrice         float64 `gorm:"check: unit_price >= 0; default: 0" json:"unit_price"`
	Discount          float64 `gorm:"check: discount >= 0; default: 0" json:"discount"`
	Tax               float64 `gorm:"check: tax >= 0; default: 0" json:"tax"`

	SaleID    uint `json:"sale_id"`
	ProductID uint `json:"product_id"`
}

type SaleIncome struct {
	ID

	SaleID   uint `json:"sale_id"`
	IncomeID uint `json:"income_id"`
}

type SaleExpense struct {
	ID

	SaleID    uint `json:"sale_id"`
	ExpenseID uint `json:"expense_id"`
}

type SaleArgument struct {
	ID

	SaleID     uint `json:"sale_id"`
	ArgumentID uint `json:"argument_id"`
}
