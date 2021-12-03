package models

import (
	"time"

	"github.com/google/uuid"
)

// Set the user status
type UserStatus struct {
	Online  bool
	Offline bool
}

// List the permissions
type PermisionList struct {
	// Dashboard
	ViewDashboard bool
	// Roles
	ViewPermissions bool
	EditPermission  bool
	ViewRoles       bool
	CreateRole      bool
	EditRole        bool
	// Users
	ViewUsers  bool
	CreateUser bool
	EditUser   bool
}

// Roles
type RoleUsers struct {
	Permissions PermisionList
}

// Set the UserInfo
type UserInfo struct {
	// User Info
	Id       uuid.UUID `json:"id" bson:"id"`
	Username string    `json:"username" bson:"username"`
	Roles    RoleUsers `json:"roles"bson:"roles"`

	// Profile Info
	ProfileImage string     `json:"profile_image"bson:"profile_image"`
	Status       UserStatus `json:"status"bson:"status"`

	// Account Info
	LastConnection time.Time `json:"last_connection"bson:"last_connection"`
	JoinedAt       time.Time `json:"joined_at"bson:"joined_at"`

	// Relations Into Accounts
	Friends       []string `json:"friends"bson:"friends"`
	Notifications []string `json:"notifications"bson:"notifications"`

	// Security
	Password     string   `json:"password"bson:"password"`
	RecoverCodes []string `json:"recover_codes"bson:"recover_codes"`
}
