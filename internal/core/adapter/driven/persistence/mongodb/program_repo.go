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

//

type ProgramRepository struct {
	program *mongo.Collection
}

const programCollections = "program"

func NewProgramRepository(db *mongo.Database) *ProgramRepository {
	col := db.Collection(programCollections)

	return &ProgramRepository{program: col}
}

func (p *ProgramRepository) AddProgram(ctx context.Context, program *education.Program) error {
	documents, err := p.getProgramDocuments(
		ctx,
		bson.M{"code": program.Code(), "specialty_id": program.SpecialtyID()},
	)
	if err != nil {
		return err
	}

	if len(documents) > 0 {
		return service.ErrProgramAlreadyExists
	}

	_, err = p.program.InsertOne(ctx, newProgramDocument(program))

	return err
}

func (p *ProgramRepository) GetProgram(ctx context.Context, id string) (*education.Program, error) {
	document, err := p.getProgramDocument(ctx, bson.M{"_id": id}, nil)
	if err != nil {
		return nil, err
	}

	return newProgram(document), nil
}

func (p *ProgramRepository) FindAllPrograms(ctx context.Context, sid string) ([]query.ProgramModel, error) {
	documents, err := p.getProgramDocuments(
		ctx,
		bson.M{"specialty_id": sid},
		options.Find().SetSort(bson.M{"code": 1}),
	)
	if err != nil {
		return []query.ProgramModel{}, err
	}

	return newProgramModels(documents), nil
}

func (p *ProgramRepository) getProgramDocument(
	ctx context.Context,
	filter bson.M,
	opts ...*options.FindOneOptions,
) (programsDocument, error) {
	var document programsDocument

	if err := p.program.FindOne(ctx, filter, opts...).Decode(&document); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return programsDocument{}, service.ErrProgramNotFound
		}

		return programsDocument{}, err
	}

	return document, nil
}

func (p *ProgramRepository) getProgramDocuments(
	ctx context.Context,
	filter bson.M,
	opts ...*options.FindOptions,
) ([]programsDocument, error) {
	var documents []programsDocument

	cursor, err := p.program.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}
