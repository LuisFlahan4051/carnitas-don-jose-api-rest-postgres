package models

import (
	"time"
)

type Sale struct {
	ID

	EntryDate   time.Time  `json:"entry_date"`
	ExitDate    *time.Time `json:"exit_date,omitempty"`
	TableNumber *uint      `json:"table_number,omitempty"`
	Packed      *bool      `json:"packed,omitempty"`

	UserID         uint           `json:"user_id"`
	BranchID       uint           `json:"branch_id"`
	TurnID         uint           `json:"turn_id"`
	SalesProducts  []SaleProduct  `json:"sales_products,omitempty"`
	SalesIncomes   []SaleIncome   `json:"sales_incomes,omitempty"`
	SalesExpenses  []SaleExpense  `json:"sales_expenses,omitempty"`
	SalesArguments []SaleArgument `json:"sales_arguments,omitempty"`
}

type SaleProduct struct {
	ID

	Done              bool     `json:"done"`
	GrammageQuantity  *uint    `json:"grammage_quantity,omitempty"`
	KilogrammagePrice *float64 `json:"kilogrammage_price,omitempty"`
	UnitQuantity      *uint    `json:"unit_quantity,omitempty"`
	UnitPrice         *float64 `json:"unit_price,omitempty"`
	Discount          *float64 `json:"discount,omitempty"`
	Tax               *float64 `json:"tax,omitempty"`

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
