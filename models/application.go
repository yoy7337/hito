package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// application column
type AppColumn struct {
	AppOId *primitive.ObjectID `bson:"appOId,omitempty" json:"appOId"` //  app._id
	App    *App                `bson:"app,omitempty" json:"app"`       // ==> App
}

// define application
type App struct {
	ID   *primitive.ObjectID `bson:"_id,omitempty" json:"_id" example:"623853b9503ce2ecdd221c94"` // ObjectId
	Name string              `bson:"name,omitempty" json:"name"`                                  // Application name
}

// implement model interface
func (_ *App) ModelName() string {
	return "apps"
}
