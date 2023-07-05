package models

import (
	"context"
	"fmt"
	"hito/utils"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const findTimeout = 5 * time.Second
const insertTimeout = 5 * time.Second
const defaultLimit = 1000
const findOneLimit int64 = 1

type FindOpts struct {
	Limit  *int64
	Skip   *int64
	Sort   interface{} // ex: {a: 1, b: -1}
	Decode interface{}
}

func (opts *FindOpts) FindOneByOId(collectionName string, oId *primitive.ObjectID) error {
	filter := bson.M{}
	if err := utils.AppendEqualTo(&filter, "_id", oId); err != nil {
		return err
	}

	return opts.FindOne(collectionName, filter)
}

func (opts *FindOpts) FindOne(collectionName string, filter interface{}) error {
	opts.Limit = utils.ToPoint(findOneLimit)
	cursor, err := doFind(collectionName, filter, opts)
	if err != nil {
		log.Fatalf("FindOne err: %s", err.Error())
		return err
	}

	if !cursor.TryNext(context.TODO()) {
		return mongo.ErrNoDocuments
	}

	if opts != nil && opts.Decode != nil {
		if err := cursor.Decode(opts.Decode); err != nil {
			log.Warnf("FindOne - decode err:  %s", err.Error())
			return err
		}
	}

	return nil
}

func (opts *FindOpts) Find(collectionName string, filter interface{}) error {
	cursor, err := doFind(collectionName, filter, opts)
	if err != nil {
		log.Warnf("Find err: %s", err.Error())
		return err
	}

	if opts.Decode != nil && cursor.TryNext(context.Background()) {
		c, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := cursor.All(c, opts.Decode); err != nil {
			log.Warnf("Find - decode err:  %s", err.Error())
			return err
		}
	}

	return nil
}

func doFind(collName string, filter interface{}, opts *FindOpts) (*mongo.Cursor, error) {
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := MongoClient.Database(DbName).Collection(collName)
	var pipeline = mongo.Pipeline{}

	if opts != nil {
	}

	// process $match
	if filter != nil {
		// { $match: { } }
		pipeline = append(pipeline, bson.D{{"$match", filter}})
	}

	// process $sort
	if opts != nil && opts.Sort != nil {
		pipeline = append(pipeline, bson.D{{"$sort", opts.Sort}})
	}
	// process $skip
	if opts != nil && opts.Skip != nil {
		pipeline = append(pipeline, bson.D{{"$skip", opts.Skip}})
	}
	// process $limit
	if opts != nil && opts.Limit != nil {
		pipeline = append(pipeline, bson.D{{"$limit", opts.Limit}})
	} else {
		pipeline = append(pipeline, bson.D{{"$limit", defaultLimit}})
	}

	log.WithFields(log.Fields{
		"pipeline": fmt.Sprintf("%+v", pipeline),
	}).Debug("doFind")

	return coll.Aggregate(c, pipeline)
}

type InsertOpts struct {
	Data     interface{}
	m        Model
	Decode   interface{}
	InsertBy *User
}

func (opts *InsertOpts) InsertOne(collName string) error {
	coll := MongoClient.Database(DbName).Collection(collName)
	c, cancel := context.WithTimeout(context.Background(), insertTimeout)
	defer cancel()

	singleResult, err := coll.InsertOne(c, opts.Data)

	log.WithFields(log.Fields{
		"singleResult": fmt.Sprintf("%+v", singleResult),
	}).Debug("InsertOne")
	if err == nil && singleResult != nil && opts != nil && opts.Decode != nil {
		//return FindOne(data.CollectionName(), bson.D{{"_id", singleResult.InsertedID}}, &FindOpts{Decode: opts.Decode})
	}

	return err
}
