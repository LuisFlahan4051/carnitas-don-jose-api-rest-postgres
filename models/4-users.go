package models

import (
	"time"
)

type Role struct {
	ID

	Name        string `gorm:"type:varchar(50)" json:"name"`
	AccessLevel uint   `gorm:"check: access_level >= 0; default: 0" json:"access_level"`

	BranchUserRoles  []BranchUserRole  `json:"branch_user_roles"`
	TurnUserRoles    []TurnUserRole    `json:"turn_user_roles"`
	InheritUserRoles []InheritUserRole `json:"inherit_user_roles"`
}

type User struct {
	ID

	Name     string `gorm:"type:varchar(50)" json:"name"`
	Lastname string `gorm:"type:varchar(50)" json:"lastname"`
	Username string `gorm:"type:varchar(50)" json:"username"`
	Password string `gorm:"type:varchar(50)" json:"password"`
	Photo    string `json:"photo"`

	Verified  bool `json:"verified"`
	Darktheme bool `json:"darktheme"`

	ActiveContract bool `json:"active_contract"`

	Address      string    `json:"address"`
	Born         time.Time `json:"born"`
	DegreeStudy  string    `gorm:"type:varchar(50)" json:"degree_study"`
	RelationShip string    `gorm:"type:varchar(50)" json:"relation_ship"`
	Curp         string    `gorm:"type:varchar(50)" json:"curp"`
	Rfc          string    `gorm:"type:varchar(50)" json:"rfc"`
	CitizenID    string    `gorm:"type:varchar(50)" json:"citizen_id"`
	CredentialID string    `gorm:"type:varchar(50)" json:"credential_id"`
	OriginState  string    `gorm:"type:varchar(50)" json:"origin_state"`

	Score          uint   `gorm:"check: score >= 0; default: 0" json:"score"`
	Qualities      string `json:"qualities"`
	Defects        string `json:"defects"`
	OriginBranchID uint   `json:"originBranch_id"`

	BranchID uint `json:"branch_id"`

	InheritUserRoles   []Role              `json:"inherit_user_roles"`
	UserPhones         []UserPhone         `json:"user_phones"`
	UserMails          []UserMail          `json:"user_mails"`
	UserReports        []UserReport        `json:"user_reports"`
	MonetaryBounds     []MonetaryBound     `json:"monetary_bounds"`
	MonetaryDiscounts  []MonetaryDiscount  `json:"monetary_discounts"`
	BranchUserRoles    []Role              `json:"branch_user_roles"`
	Turns              []Turn              `json:"turns"`
	TurnUserRoles      []TurnUserRole      `json:"turn_user_roles"`
	Sales              []Sale              `json:"sales"`
	UserSafeboxActions []UserSafeboxAction `json:"safebox_actions"`
	AdminNotifications []AdminNotification `json:"admin_notifications"`
}

type InheritUserRole struct {
	ID

	RoleID uint `json:"role_id"`
	UserID uint `json:"user_id"`
}

type UserPhone struct {
	ID

	Phone string `gorm:"type:varchar(50)" json:"phone"`
	Main  bool   `json:"main"`

	UserID uint `json:"user_id"`
}

type UserMail struct {
	ID

	Mail string `gorm:"type:varchar(50)" json:"mail"`
	Main bool   `json:"main"`

	UserID uint `json:"user_id"`
}

type UserReport struct {
	ID

	Reason string `json:"reason"`

	UserID uint `json:"user_id"`
}

type MonetaryBound struct {
	ID

	Reason string  `json:"reason"`
	Bound  float64 `json:"bound"`

	UserID uint `json:"user_id"`
}

type MonetaryDiscount struct {
	ID

	Reason   string  `json:"reason"`
	Discount float64 `json:"discount"`

	UserID uint `json:"user_id"`
}

type BranchUserRole struct {
	ID

	BranchID uint `json:"branch_id"`
	UserID   uint `json:"user_id"`
	RoleID   uint `json:"role_id"`
}
