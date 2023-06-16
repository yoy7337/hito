package models

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"hito/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserColumn struct {
	UserOId *primitive.ObjectID `bson:"userOId,omitempty" json:"userOId"` //  user._id
	User    *User               `bson:"user,omitempty" json:"user"`       // ==> user
}

type User struct {
	BaseModel    `bson:"inline"`
	Id           string `bson:"id,omitempty" json:"id"`                                    // id for login
	Name         string `bson:"name,omitempty" json:"name"`                                // user name
	Email        string `bson:"email,omitempty" json:"email" example:"hito@hito.com"`      // email
	PhoneNo      string `bson:"phoneNo,omitempty" json:"phoneNo" example:"+886 912345678"` // phone number
	PasswordSalt string `bson:"_password_salt,omitempty" json:"-"`
	PasswordHash string `bson:"_password_hash,omitempty" json:"-"`
}

// implement model interface
func (_ *User) ModelName() string {
	return "users"
}

func (u *User) Authentication(pass string) bool {
	return u.PasswordHash == u.saltHashPassword(pass, u.PasswordSalt)
}

func (u *User) SetPassword(pass string) error {
	salt, err := utils.GenerateRandomHexString(16)
	if err != nil {
		return err
	}

	u.PasswordSalt = salt
	u.PasswordHash = u.saltHashPassword(pass, salt)

	return nil
}

func (u *User) saltHashPassword(pass string, salt string) string {
	h := hmac.New(sha512.New, []byte(salt))
	h.Write([]byte(pass))

	return hex.EncodeToString(h.Sum(nil))
}
