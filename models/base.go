package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ModelStatus int8

const (
	Normal ModelStatus = iota
	Disable
	Delete uint8 = 9
)

type BaseModel struct {
	AppColumn `bson:"inline"`
	ID        *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c94"`             // ObjectId
	Status    *ModelStatus        `bson:"status,omitempty" json:"status" example:"0"`                              // model status
	Sort      *uint64             `bson:"sort,omitempty" json:"sort" example:"0"`                                  // sorting priority
	CreatedAt *time.Time          `bson:"createdAt,omitempty" json:"createdAt" example:"2022-03-21T10:30:17.711Z"` // create timestamp
	UpdatedAt *time.Time          `bson:"updatedAt,omitempty" json:"updatedAt" example:"2022-03-21T10:30:17.711Z"` // update timestamp
	CreatedBy *UserColumn         `bson:"createdBy,omitempty" json:"createdBy" example:"623853b9503ce2ecdd221c94"` // created by (linked to user)
	UpdatedBy *UserColumn         `bson:"updatedBy,omitempty" json:"updatedBy" example:"623853b9503ce2ecdd221c94"` // updated by (linked to user)
}

// interface for general model
type Model interface {
	ModelName() string // return model(collection) name
}
