package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserColumn struct {
	UserOId *primitive.ObjectID `bson:"userOId,omitempty" json:"userOId"` //  user._id
	User    *User               `bson:"user,omitempty" json:"user"`       // ==> user
}

type User struct {
	BaseModel    `bson:"inline"`
	LoginId      string `bson:"loginId,omitempty" json:"loginId"`                          // id for login
	Name         string `bson:"name,omitempty" json:"name"`                                // user name
	PhoneNo      string `bson:"phoneNo,omitempty" json:"phoneNo" example:"+886 912345678"` // phone number
	PasswordSalt string `bson:"_password_salt,omitempty" json:"-"`
	PasswordHash string `bson:"_password_hash,omitempty" json:"-"`
}

// implement model interface
func (_ *User) ModelName() string {
	return "users"
}
