package mongo

import (
	"context"
	"fmt"
	"log"
	"strconv"

	beer "github.com/ericuss/golangSimpleKatas/pkg"
	"go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type beerRepository struct {
	beers      map[int]*beer.Beer
	client     *mongo.Client
	collection *mongo.Collection
}

func NewBeerRepository(beers map[int]*beer.Beer) beer.BeerRepository {
	// host := "localhost"
	host := "mongo"
	port := 27017

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Congratulations, you're already connected to MongoDB!")
	collection := client.Database("Brewery").Collection("Beers")

	return &beerRepository{
		beers:      beers,
		client:     client,
		collection: collection,
	}
}

func (r *beerRepository) Fetch() ([]*beer.Beer, error) {
	var results []*beer.Beer
	findOptions := options.Find()
	cur, err := r.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var s beer.Beer
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &s)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Beers found: %v\n", results)

	return results, nil
}

func (r *beerRepository) FetchById(id int) (res *beer.Beer, err error) {
	// var result *beer.Beer
	// filter := bson.D{"name", id}
	// err = r.collection.FindOne(context.TODO(), filter).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil, fmt.Errorf("The Id %s doesn't exist", strconv.Itoa(id))
	// }
	// fmt.Printf("Beers found: %v\n", result)
	// return result, nil
	return nil, fmt.Errorf("The Id %s doesn't exist", strconv.Itoa(id))

}
