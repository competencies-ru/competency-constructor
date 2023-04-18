package mongodb

import (
	"context"
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

const subjectCollections = "subject"

type SubjectRepository struct {
	subjects *mongo.Collection
}

func NewSubjectRepository(db *mongo.Database) *SubjectRepository {
	col := db.Collection(subjectCollections)

	_, err := col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		panic(err)
	}

	return &SubjectRepository{subjects: col}
}

func (r *SubjectRepository) AddSubject(ctx context.Context, subject *competencies.Subject) error {
	_, err := r.subjects.InsertOne(ctx, newSubjectDocument(subject))

	return err
}

func (r *SubjectRepository) AddSubjects(ctx context.Context, subjects []*competencies.Subject) error {
	_, err := r.subjects.InsertMany(ctx, newSubjectDocuments(subjects))

	return err
}

func (r *SubjectRepository) FilterSubjectsByName(ctx context.Context, name string) []query.SubjectModel {
	documents, err := r.getSubjectDocuments(ctx, makeFilterSubjectsByName(name))
	if err != nil {
		return nil
	}

	return newSubjectModels(documents)
}

func (r *SubjectRepository) getSubjectDocuments(
	ctx context.Context,
	filter interface{},
	opts ...*options.FindOptions,
) ([]subjectDocument, error) {
	var documents []subjectDocument

	cursor, err := r.subjects.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	return documents, nil
}

func makeFilterSubjectsByName(name string) bson.D {
	builder := strings.Builder{}
	builder.WriteString("/.*")
	builder.WriteString(name)
	builder.WriteString("*/")

	return bson.D{{"name", bson.D{{"$regex", builder.String()}}}}
}
