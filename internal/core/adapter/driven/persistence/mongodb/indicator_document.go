package mongodb

type (
	indicatorDocument struct {
		ID           string `bson:"_id"`
		Title        string `bson:"title"`
		Code         string `bson:"code"`
		SubjectID    string `bson:"subjectID"`
		CompetencyID string `bson:"competencyID"`
	}
)
