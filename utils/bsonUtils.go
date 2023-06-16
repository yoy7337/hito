package utils

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StructToBson(src interface{}) (*bson.M, error) {
	bytes, err := bson.Marshal(src)
	if err != nil {
		return nil, err
	}

	bsonData := bson.M{}
	if err = bson.Unmarshal(bytes, bsonData); err != nil {
		return nil, err
	}

	return &bsonData, nil
}

func BsonToInterface(src interface{}) map[string]interface{} {
	doc, err := bson.Marshal(src)
	if err != nil {
		return nil
	}

	dest := map[string]interface{}{}
	if err = bson.Unmarshal(doc, &dest); err != nil {
		return nil
	}

	return dest
}

func AppendObjId(bsonMap *bson.M, key string, val string) error {
	if bsonMap == nil {
		return errors.New("bsonMap is nil")
	}

	objId, err := primitive.ObjectIDFromHex(val)
	if err != nil {
		return err
	}

	(*bsonMap)[key] = objId

	return nil
}

func AppendEqualTo(bsonMap *bson.M, key string, val interface{}) error {
	if bsonMap == nil {
		return errors.New("bsonMap is nil")
	}

	(*bsonMap)[key] = val

	return nil
}

func AppendInQuery[T any](bsonMap *bson.M, key string, arr *[]T) error {
	if bsonMap == nil {
		return errors.New("bsonMap is nil")
	}

	if arr != nil && len(*arr) > 0 {
		bsonA := bson.A{}
		for _, val := range *arr {
			bsonA = append(bsonA, val)
		}

		(*bsonMap)[key] = bson.D{{"$in", bsonA}}
	}

	return nil
}

func AppendRegexQuery(bsonMap *bson.M, key string, val string) {
	(*bsonMap)[key] = bson.D{{"$regex", primitive.Regex{Pattern: val, Options: "im"}}}
}
