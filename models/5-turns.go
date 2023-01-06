package models

import (
	"time"
)

type Turn struct {
	ID

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Active    bool      `json:"active"`

	IncomesCounter float64 `json:"incomes_counter"`
	NetIncomesCounter float64 `json:"netincomes_counter"`
	ExpensesCounter float64 `json:"expenses_counter"`


	UserID        uint           `json:"user_id"`
	BranchID      uint           `json:"branch_id"`
	TurnUserRoles []TurnUserRole `json:"user_roles"`
	SafeboxReceived TurnSafebox   `json:"safebox_received"`
	SafeboxFinished TurnSafebox `json:"safebox_finished"`
	Sales         []Sale      `json:"sales"`
	Inventories   []Inventory `json:"inventories"`
}

type TurnUserRole struct {
	ID

	LoginDate  time.Time `json:"login_date"`
	LogoutDate time.Time `json:"logout_date"`

	UserID uint `json:"user_id"`
	TurnID uint `json:"turn_id"`
	RoleID uint `json:"role_id"`
}

type TurnSafebox struct {
	ID

	TurnID    uint `json:"turn_id"`
	SafeboxID uint `gorm:"unique" json:"safebox_id"`
}
