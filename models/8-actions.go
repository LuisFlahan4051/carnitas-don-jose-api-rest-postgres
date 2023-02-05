package models

import "time"

type UserSafeboxAction struct {
	ID

	Withdrawal bool `json:"withdrawal"`

	UserID        uint `json:"user_id"`
	ActionSafebox `json:"action_safebox"`
}

type ActionSafebox struct {
	ID

	SafeboxActionID uint `gorm:"unique" json:"safebox_action_id"`
	SafeboxID       uint `gorm:"unique" json:"safebox_id"`
}

type AdminNotification struct {
	ID

	Type        string `json:"type"`
	Solved      bool   `json:"solved"`
	Description string `json:"description"`

	BranchID *uint                    `json:"branch_id,omitempty"`
	UserID   *uint                    `json:"user_id,omitempty"`
	Images   []AdminNotificationImage `json:"images,omitempty"`
}

type AdminNotificationImage struct {
	ID

	Image string `json:"image"`

	NotificationID uint `json:"notification_id"`
}

type ServerLogs struct {
	Id          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Transaction string    `json:"transaction"`
	UserID      uint      `json:"user_id"`
	BranchID    *uint     `json:"branch_id,omitempty"`
	Root        *bool     `json:"root,omitempty"`
}

type Pagination struct {
	Page  *int       `json:"page,omitempty"`
	Today *bool      `json:"today"`
	Since *time.Time `json:"since"`
	To    *time.Time `json:"to"`
}
