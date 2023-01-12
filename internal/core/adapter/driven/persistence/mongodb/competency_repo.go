package mongodb

import (
	"context"

	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CompetencyRepository struct {
	competency *mongo.Collection
}

const competenciesCollection = "competencies"

func NewCompetencyRepository(db *mongo.Database) *CompetencyRepository {
	col := db.Collection(competenciesCollection)

	return &CompetencyRepository{competency: col}
}

func (r *CompetencyRepository) AddCompetency(ctx context.Context, competency *competencies.Competency) error {
	documents, err := r.geCompetencyDocuments(ctx, makeFilterCompetencyOnType(competency))
	if err != nil {
		return err
	}

	if len(documents) > 0 {
		return service.ErrCompetencyAlreadyExists
	}

	_, err = r.competency.InsertOne(ctx, newCompetencyDocument(competency))

	return err
}

func (r *CompetencyRepository) GetCompetency(ctx context.Context, id string) (*competencies.Competency, error) {
	document, err := r.geCompetencyDocument(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return newCompetency(document), nil
}

func (r *CompetencyRepository) FilterCompetencies(
	ctx context.Context,
	levelID,
	ugsnID,
	specialtyID,
	programID string,
) ([]query.CompetencyModel, error) {

	sort := options.Find().SetSort(bson.M{"code": 1})

	documents, err := r.geCompetencyDocuments(ctx, makeFilterCompetency(levelID, ugsnID, specialtyID, programID), sort)
	if err != nil {
		return nil, err
	}

	return newCompetencyModels(documents), nil
}

func makeFilterCompetencyOnType(competency *competencies.Competency) bson.M {
	switch competency.CompetencyType() {
	case competencies.UNIVERSAL:
		return bson.M{"type": competency.CompetencyType(), "level_id": competency.LevelID(), "code": competency.Code()}
	case competencies.PROFESSIONAL:
		return bson.M{"type": competency.CompetencyType(), "program_id": competency.ProgramID(), "code": competency.Code()}
	case competencies.GENERAL:
		if competency.UgsnID() != "" {
			return bson.M{"type": competency.CompetencyType(), "ugsn_id": competency.UgsnID(), "code": competency.Code()}
		}

		return bson.M{"type": competency.CompetencyType(), "specialty_id": competency.SpecialtyID(), "code": competency.Code()}
	}

	return bson.M{}
}

func makeFilterCompetency(levelID, ugsnID, specialtyID, programID string) bson.D {
	m := make(map[string]interface{})
	if levelID != "" {
		m["level_id"] = levelID
	}

	if ugsnID == "" {
		m["ugsn_id"] = ugsnID
	}

	if specialtyID == "" {
		m["specialty_id"] = specialtyID
	}

	if programID == "" {
		m["program_id"] = programID
	}

	return bson.D{{Key: "$or", Value: bson.A{m}}}
}

func (r *CompetencyRepository) geCompetencyDocument(
	ctx context.Context,
	filter bson.M,
	opts ...*options.FindOneOptions,
) (competencyDocument, error) {
	var document competencyDocument

	if err := r.competency.FindOne(ctx, filter, opts...).Decode(&document); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return competencyDocument{}, service.ErrSpecialtyNotFound
		}

		return competencyDocument{}, err
	}

	return document, nil
}

func (r *CompetencyRepository) geCompetencyDocuments(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOptions,
) ([]competencyDocument, error) {
	var documents []competencyDocument

	cursor, err := r.competency.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}
