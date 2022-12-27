package mongodb

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LevelRepository struct {
	level *mongo.Collection
}

const levelCollections = "levels"

func NewLevelRepository(db *mongo.Database) *LevelRepository {
	col := db.Collection(levelCollections)

	_, err := col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"title": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}

	return &LevelRepository{
		level: col,
	}
}

func (r *LevelRepository) AddLevel(ctx context.Context, level *education.Level) error {
	_, err := r.level.InsertOne(ctx, newLevelDocument(level))

	return err
}

func (r *LevelRepository) GetLevel(ctx context.Context, id string) (*education.Level, error) {
	// TODO implement me
	panic("implement me")
}

func (r *LevelRepository) FindLevel(ctx context.Context, id string) (query.SpecificLevelModel, error) {
	filter := bson.M{"_id": id}

	document, err := r.getLevelDocument(ctx, filter, nil)
	if err != nil {
		return query.SpecificLevelModel{}, err
	}

	return newSpecificLevelView(document), nil
}

func (r *LevelRepository) FindLevels(ctx context.Context) ([]query.LevelModel, error) {
	var documents []levelDocument

	opts := options.Find().SetProjection(bson.M{"_id": 1, "title": 1})

	cursor, err := r.level.Find(ctx, bson.D{}, opts)
	if err != nil {
		return []query.LevelModel{}, err
	}

	if err := cursor.All(ctx, &documents); err != nil {
		return []query.LevelModel{}, err
	}

	return newLevelModelView(documents), err
}

func (r *LevelRepository) UpdateLevel(ctx context.Context, id string, updater service.LevelUpdate) error {
	session, err := r.level.Database().Client().StartSession()
	if err != nil {
		return err
	}

	defer session.EndSession(ctx)

	_, err = session.WithTransaction(ctx, func(_ mongo.SessionContext) (interface{}, error) {
		document, err := r.getLevelDocument(ctx, bson.M{"_id": id}, nil)
		if err != nil {
			return nil, err
		}

		level := newLevel(document)

		updaterUgsn, err := updater(ctx, level)
		if err != nil {
			return nil, err
		}

		updatedDocument := newLevelDocument(updaterUgsn)

		replaceOpt := options.Replace().SetUpsert(true)
		filter := bson.M{"_id": updatedDocument.ID}
		if _, err := r.level.ReplaceOne(ctx, filter, updatedDocument, replaceOpt); err != nil {
			return nil, err
		}

		return nil, nil
	})

	return err
}

func (r *LevelRepository) getLevelDocument(
	ctx context.Context,
	filter bson.M,
	opts ...*options.FindOneOptions,
) (levelDocument, error) {
	var document levelDocument

	if err := r.level.FindOne(ctx, filter, opts...).Decode(&document); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return levelDocument{}, service.ErrLevelNotFound
		}

		return levelDocument{}, err
	}

	return document, nil
}
