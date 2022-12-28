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

	BranchID uint                     `json:"branch_id"`
	UserID   uint                     `json:"turn_id"`
	Images   []AdminNotificationImage `json:"images"`
}

type AdminNotificationImage struct {
	ID

	Image string `json:"image"`

	NotificationID uint `json:"notification_id"`
}

type ServerLogs struct {
	Id          uint      `json:"id"`
	CreateAt    time.Time `json:"create_at"`
	Transaction string    `json:"transaction"`
	UserID      uint      `json:"user_id"`
	BranchID    *uint     `json:"branch_id"`
	Root        *bool     `json:"root"`
}
