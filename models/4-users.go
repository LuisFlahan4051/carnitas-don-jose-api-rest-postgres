package models

import (
	"time"
)

type Role struct {
	ID

	Name        string `json:"name"`
	AccessLevel uint   `json:"access_level"`

	BranchUserRoles  []BranchUserRole  `json:"branch_user_roles,omitempty"`
	TurnUserRoles    []TurnUserRole    `json:"turn_user_roles,omitempty"`
	InheritUserRoles []InheritUserRole `json:"inherit_user_roles,omitempty"`
}

type User struct {
	ID

	Name     *string `json:"name,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Username string  `json:"username"`
	Password string  `json:"password,omitempty"`
	Photo    *string `json:"photo,omitempty"`

	// Just need to get the bytes of the file and save it.
	// See registUser() at github.com/LuisFlahan4051/carnitas-don-jose-api-rest-postgres/routes/supervisorActions/actions.go
	//ProfilePicture *multipart.File `json:"profile_picture,omitempty"`

	Verified  *bool `json:"verified,omitempty"`
	Warning   *bool `json:"warning,omitempty"`
	Darktheme *bool `json:"darktheme,omitempty"`

	ActiveContract *bool `json:"active_contract,omitempty"`

	Address      *string    `json:"address,omitempty"`
	Born         *time.Time `json:"born,omitempty"`
	DegreeStudy  *string    `json:"degree_study,omitempty"`
	RelationShip *string    `json:"relation_ship,omitempty"`
	Curp         *string    `json:"curp,omitempty"`
	Rfc          *string    `json:"rfc,omitempty"`
	CitizenID    *string    `json:"citizen_id,omitempty"`
	CredentialID *string    `json:"credential_id,omitempty"`
	OriginState  *string    `json:"origin_state,omitempty"`

	Score          *uint   `json:"score,omitempty"`
	Qualities      *string `json:"qualities,omitempty"`
	Defects        *string `json:"defects,omitempty"`
	OriginBranchID *uint   `json:"origin_branch_id,omitempty"`

	BranchID *uint `json:"branch_id,omitempty"`

	InheritUserRoles   []Role              `json:"inherit_user_roles,omitempty"`
	UserPhones         []UserPhone         `json:"user_phones,omitempty"`
	UserMails          []UserMail          `json:"user_mails,omitempty"`
	UserReports        []UserReport        `json:"user_reports,omitempty"`
	MonetaryBounds     []MonetaryBound     `json:"monetary_bounds,omitempty"`
	MonetaryDiscounts  []MonetaryDiscount  `json:"monetary_discounts,omitempty"`
	BranchUserRoles    []Role              `json:"branch_user_roles,omitempty"`
	Turns              []Turn              `json:"turns,omitempty"`
	TurnUserRoles      []TurnUserRole      `json:"turn_user_roles,omitempty"`
	Sales              []Sale              `json:"sales,omitempty"`
	UserSafeboxActions []UserSafeboxAction `json:"safebox_actions,omitempty"`
	AdminNotifications []AdminNotification `json:"admin_notifications,omitempty"`
}

type InheritUserRole struct {
	ID

	RoleID uint `json:"role_id"`
	UserID uint `json:"user_id"`
}

type UserPhone struct {
	ID

	Phone string `json:"phone"`
	Main  *bool  `json:"main,omitempty"`

	UserID uint `json:"user_id"`
}

type UserMail struct {
	ID

	Mail string `json:"mail"`
	Main *bool  `json:"main,omitempty"`

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
