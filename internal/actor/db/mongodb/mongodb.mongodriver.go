package mongodb

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongodb Implementation
//

type mongodb struct {
	uri         string
	name        string
	retryWrites bool
	client      *mongo.Client
	logger      zerolog.Logger
}

func NewMongodb(uri string, name string, retryWrites bool) Mongodb {
	return &mongodb{
		uri:         uri,
		name:        name,
		retryWrites: retryWrites,
		logger: log.
			With().
			Str("actor", "db/mongodb").
			Logger(),
	}
}

func (r *mongodb) Connect() Mongodb {
	r.logger.
		Info().
		Str("dbName", r.name).
		Str("status", "connecting...").
		Send()

	if r.name == "" {
		r.logger.
			Fatal().
			Str("status", "missing parameters").
			Str("reason", "name is required").
			Send()
	}
	if r.uri == "" {
		r.logger.
			Fatal().
			Str("dbName", r.name).
			Str("status", "missing parameters").
			Str("reason", "uri is required").
			Send()
	}
	connectionURI := r.uri
	// if r.retryWrites {
	// 	connectionURI = connectionURI + "/?retryWrites=true"
	// }

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		r.logger.
			Fatal().
			Err(err).
			Str("dbName", r.name).
			Str("status", "connection failed!").
			Send()
	}
	r.client = client
	err = r.client.Ping(ctx, nil)
	if err != nil {
		r.logger.
			Fatal().
			Err(err).
			Str("dbName", r.name).
			Str("status", "failed to ping client").
			Send()
	}
	r.logger.
		Info().
		Str("dbName", r.name).
		Str("status", "connected successfully.").
		Send()
	return r
}

func (r *mongodb) Disconnect() {
	r.logger.
		Info().
		Str("dbName", r.name).
		Str("status", "disconnecting...").
		Send()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := r.client.Disconnect(ctx)
	if err != nil {
		r.logger.
			Debug().
			Err(err).
			Str("dbName", r.name).
			Str("status", "failed to disconnect gracefully").
			Send()
	}
	r.logger.
		Info().
		Str("dbName", r.name).
		Str("status", "disconnected successfully").
		Send()
}

func (r *mongodb) GetDatabase() Database {
	return NewDatabase(r.client.Database(r.name))
}

func (db *mongodb) Collection(name string, opts ...*options.CollectionOptions) Collection {
	mongoColl := db.client.Database(db.name).Collection(name, opts...)
	return NewCollection(mongoColl)
}

func (r *mongodb) Ping() error {
	return r.client.Ping(context.TODO(), nil)
}

// Database Implementation
//

type database struct {
	db *mongo.Database
}

func NewDatabase(mongoDB *mongo.Database) Database {
	return &database{db: mongoDB}
}

func (d *database) Client() *mongo.Client {
	return d.db.Client()
}

func (d *database) Name() string {
	return d.db.Name()
}

func (d *database) Collection(name string, opts ...*options.CollectionOptions) Collection {
	return &collection{d.db.Collection(name, opts...)}
}

func (d *database) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return d.db.Aggregate(ctx, pipeline, opts...)
}

func (d *database) RunCommand(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) *mongo.SingleResult {
	return d.db.RunCommand(ctx, runCommand, opts...)
}

func (d *database) RunCommandCursor(ctx context.Context, runCommand interface{}, opts ...*options.RunCmdOptions) (*mongo.Cursor, error) {
	return d.db.RunCommandCursor(ctx, runCommand, opts...)
}

func (d *database) Drop(ctx context.Context) error {
	return d.db.Drop(ctx)
}

func (d *database) ListCollectionSpecifications(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]*mongo.CollectionSpecification, error) {
	return d.db.ListCollectionSpecifications(ctx, filter, opts...)
}

func (d *database) ListCollections(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) (*mongo.Cursor, error) {
	return d.db.ListCollections(ctx, filter, opts...)
}

func (d *database) ListCollectionNames(ctx context.Context, filter interface{}, opts ...*options.ListCollectionsOptions) ([]string, error) {
	return d.db.ListCollectionNames(ctx, filter, opts...)
}

func (d *database) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return d.db.Watch(ctx, pipeline, opts...)
}

func (d *database) CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error {
	return d.db.CreateCollection(ctx, name, opts...)
}

func (d *database) CreateView(ctx context.Context, viewName, viewOn string, pipeline interface{}, opts ...*options.CreateViewOptions) error {
	return d.db.CreateView(ctx, viewName, viewOn, pipeline, opts...)
}

// Collection Implementation
//

type collection struct {
	collection *mongo.Collection
}

func NewCollection(coll *mongo.Collection) Collection {
	return &collection{collection: coll}
}

func (c *collection) Name() string {
	return c.collection.Name()
}
func (c *collection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	return c.collection.BulkWrite(ctx, models, opts...)
}

func (c *collection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.collection.InsertOne(ctx, document, opts...)
}

func (c *collection) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return c.collection.InsertMany(ctx, documents, opts...)
}

func (c *collection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return c.collection.DeleteOne(ctx, filter, opts...)
}

func (c *collection) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return c.collection.DeleteMany(ctx, filter, opts...)
}

func (c *collection) UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	filter := bson.M{"_id": id}
	return c.collection.UpdateOne(ctx, filter, bson.M{"$set": update}, opts...)
}

func (c *collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.collection.UpdateOne(ctx, filter, update, opts...)
}

func (c *collection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return c.collection.UpdateMany(ctx, filter, update, opts...)
}

func (c *collection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return c.collection.ReplaceOne(ctx, filter, replacement, opts...)
}

func (c *collection) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return c.collection.Aggregate(ctx, pipeline, opts...)
}

func (c *collection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return c.collection.CountDocuments(ctx, filter, opts...)
}

func (c *collection) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return c.collection.EstimatedDocumentCount(ctx, opts...)
}

func (c *collection) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return c.collection.Distinct(ctx, fieldName, filter, opts...)
}

func (c *collection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return c.collection.Find(ctx, filter, opts...)
}

func (c *collection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return c.collection.FindOne(ctx, filter, opts...)
}
