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

type UgsnRepository struct {
	ugsn *mongo.Collection
}

const ugsnCollection = "ugsn"

func NewUgsnRepository(db *mongo.Database) *UgsnRepository {
	col := db.Collection(ugsnCollection)

	return &UgsnRepository{ugsn: col}
}

func (u *UgsnRepository) FindAllUgsn(ctx context.Context, levelID string) ([]query.UgsnModel, error) {
	documents, err := u.geUgsnDocuments(
		ctx,
		bson.M{"level_id": levelID},
		options.Find().SetSort(bson.M{"code": 1}),
	)
	if err != nil {
		return nil, err
	}

	return newUgsnModels(documents), err
}

func (u *UgsnRepository) AddUgsn(ctx context.Context, ugsn *education.Ugsn) error {
	documents, err := u.geUgsnDocuments(ctx, bson.M{"code": ugsn.Code(), "level_id": ugsn.LeveID()})
	if err != nil {
		return err
	}

	if len(documents) > 0 {
		return service.ErrUgsnAlreadyExists
	}

	_, err = u.ugsn.InsertOne(ctx, newUgsnDocument(ugsn))

	return err
}

func (u *UgsnRepository) GetUgsn(ctx context.Context, id string) (*education.Ugsn, error) {
	document, err := u.geUgsnDocument(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return newUgsn(document), nil
}

func (u *UgsnRepository) geUgsnDocument(
	ctx context.Context,
	filter bson.M,
	opts ...*options.FindOneOptions,
) (ugsnDocument, error) {
	var document ugsnDocument

	if err := u.ugsn.FindOne(ctx, filter, opts...).Decode(&document); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ugsnDocument{}, service.ErrUgsnNotFound
		}

		return ugsnDocument{}, err
	}

	return document, nil
}

func (u *UgsnRepository) geUgsnDocuments(
	ctx context.Context,
	filter bson.M,
	opts ...*options.FindOptions,
) ([]ugsnDocument, error) {
	var documents []ugsnDocument

	cursor, err := u.ugsn.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}
