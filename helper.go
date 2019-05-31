package mongohelper

import "go.mongodb.org/mongo-driver/bson"
import "go.mongodb.org/mongo-driver/bson/primitive"
import "go.mongodb.org/mongo-driver/mongo"
import "go.mongodb.org/mongo-driver/mongo/options"
import "context"

//TextSearchFilter returns a text search filter
func TextSearchFilter(search string) bson.M {
	filter := bson.M{}
	if search != "" {
		filter = bson.M{
			"$text": bson.M{
				"$search": search,
			},
		}
	}
	return filter
}

//ObjectIdFilter returns a ObjectId filter
func ObjectIdFilter(id string) (bson.M, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return bson.M{"_id": oid}, err
}

//FindOnByObjectId finds one by _id
func FindOneByObjectId(coll *mongo.Collection, ctx context.Context, objectId string,
	opts ...*options.FindOneOptions) (*mongo.SingleResult, error) {
	filter, err := ObjectIdFilter(objectId)
	if err != nil {
		return nil, err
	}
	r := coll.FindOne(ctx, filter, opts...)
	return r, err
}

//FindByTextSearch finds many by a text search
func FindByTextSearch(coll *mongo.Collection, ctx context.Context, search string,
	opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return coll.Find(ctx, TextSearchFilter(search), opts...)
}
