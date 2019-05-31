package mongohelper

import "go.mongodb.org/mongo-driver/bson"
import "go.mongodb.org/mongo-driver/bson/primitive"

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
