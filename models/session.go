package models

import "time"

// define a session model
type Session struct {
	BaseModel    `bson:"inline"`
	UserColumn   `bson:"inline"`
	SessionToken string     `bson:"sessionToken,omitempty"`
	ExpiresAt    *time.Time `bson:"expiresAt,omitempty"`
}

// implement model interface
func (_ *Session) ModelName() string {
	return "sessions"
}
