package crud

import (
	"context"
    "kano-library/mongo"
	"kano-library/util"
	"kano-library/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(collection string, query models.Books) (string, error) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)
	query.Id = primitive.NewObjectID()

	_, err = db.Collection(collection).InsertOne(ctx, query)
	util.ErrorChecker(err)
	return "Inserted Successfully", err
}
