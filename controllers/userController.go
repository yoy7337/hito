package controllers

import (
	"errors"
	"fmt"
	"hito/models"
	"hito/utils"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type SignUpOpts struct {
	Id       string `json:"id" binding:"required" example:"user1"`
	Password string `json:"password" binding:"required" example:"123456"`
	Email    string `json:"email" binding:"required,email" example:"hito@hito.com"`
	Name     string `json:"name" example:"hitooooo"`
}

func (opts *SignUpOpts) SignUp(result *models.User) error {
	log.WithFields(log.Fields{
		"opts": fmt.Sprintf("%+v", *opts),
	}).Debug("opts.signUp")

	// check if user exists
	var users []models.User
	findOpts := models.FindOpts{
		Decode: &users,
	}

	filter := bson.M{}
	if err := utils.AppendEqualTo(&filter, "id", opts.Id); err != nil {
		return err
	}

	if err := findOpts.Find("users", filter); err != nil {
		return err
	}

	if users != nil {
		log.WithFields(log.Fields{
			"users": fmt.Sprintf("%+v", users),
		}).Debug("opts.signUp.findUsers")

		return errors.New("user already exists")
	}

	// create user
	user := models.User{
		Id:    opts.Id,
		Name:  opts.Name,
		Email: opts.Email,
	}
	// setup password
	if err := user.SetPassword(opts.Password); err != nil {
		return err
	}

	// insert user
	insertOpts := models.InsertOpts{
		Data:   user,
		Decode: result,
	}

	return insertOpts.InsertOne("users")
}

type LoginOpts struct {
	Id       string `json:"id" binding:"required" example:"user1"`
	Password string `json:"password" binding:"required" example:"123456"`
}

func (opts *LoginOpts) Login(user *models.User) error {
	// check if user exists
	findOpts := models.FindOpts{
		Decode: user,
	}

	filter := bson.M{}
	if err := utils.AppendEqualTo(&filter, "id", opts.Id); err != nil {
		return err
	}

	if err := findOpts.FindOne("users", filter); err != nil {
		return err
	}

	if ok := user.Authentication(opts.Password); ok == false {
		return errors.New("authentication failed")
	}

	return nil
}
