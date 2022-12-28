package models

import (
	"time"
)

type Turn struct {
	ID

	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Active    bool      `json:"active"`

	UserID        uint           `json:"user_id"`
	BranchID      uint           `json:"branch_id"`
	TurnUserRoles []TurnUserRole `json:"turn_user_roles"`
	TurnSafebox   `json:"turn_safebox"`
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

	TurnID    uint `gorm:"unique" json:"turn_id"`
	SafeboxID uint `gorm:"unique" json:"safebox_id"`
}
