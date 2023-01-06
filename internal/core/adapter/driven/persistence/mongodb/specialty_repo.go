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

type SpecialtyRepository struct {
	specialty *mongo.Collection
}

const specialtyCollections = "specialty"

func NewSpecialtyRepository(db *mongo.Database) *SpecialtyRepository {
	col := db.Collection(specialtyCollections)

	return &SpecialtyRepository{specialty: col}
}

func (s *SpecialtyRepository) FindAllSpecialties(ctx context.Context, uid string) ([]query.SpecialtyModel, error) {
	documents, err := s.geSpecialtyDocuments(
		ctx,
		bson.M{"ugsn_id": uid},
		options.Find().SetSort(bson.M{"code": 1}),
	)
	if err != nil {
		return nil, err
	}

	return newSpecialtyModels(documents), nil
}

func (s *SpecialtyRepository) AddSpecialty(ctx context.Context, specialty *education.Speciality) (err error) {
	documents, err := s.geSpecialtyDocuments(
		ctx,
		bson.M{"code": specialty.Code(), "ugsn_id": specialty.UgsnID()},
	)
	if err != nil {
		return err
	}

	if len(documents) > 0 {
		return service.ErrSpecialtyAlreadyExists
	}

	_, err = s.specialty.InsertOne(ctx, newSpecialtyDocument(specialty))

	return err
}

func (s *SpecialtyRepository) GetSpecialty(ctx context.Context, id string) (*education.Speciality, error) {
	document, err := s.geSpecialtyDocument(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return newSpecialty(document), nil
}

func (s *SpecialtyRepository) geSpecialtyDocument(
	ctx context.Context,
	filter bson.M,
	opts ...*options.FindOneOptions,
) (specialtiesDocument, error) {
	var document specialtiesDocument

	if err := s.specialty.FindOne(ctx, filter, opts...).Decode(&document); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return specialtiesDocument{}, service.ErrSpecialtyNotFound
		}

		return specialtiesDocument{}, err
	}

	return document, nil
}

func (s *SpecialtyRepository) geSpecialtyDocuments(
	ctx context.Context,
	filter bson.M,
	opts ...*options.FindOptions,
) ([]specialtiesDocument, error) {
	var documents []specialtiesDocument

	cursor, err := s.specialty.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}
