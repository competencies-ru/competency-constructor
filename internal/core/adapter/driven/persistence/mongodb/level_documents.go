package mongodb

import (
	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/education"
)

type (
	levelDocument struct {
		ID    string         `bson:"_id,omitempty"`
		Title string         `bson:"title,omitempty"`
		Ugsn  []ugsnDocument `bson:"ugsn"`
	}

	ugsnDocument struct {
		ID          string                `bson:"id,omitempty"`
		Code        string                `bson:"code"`
		Title       string                `bson:"title"`
		Specialties []specialtiesDocument `bson:"specialties"`
	}

	specialtiesDocument struct {
		ID       string             `bson:"id,omitempty"`
		Code     string             `bson:"code"`
		Title    string             `bson:"title"`
		Programs []programsDocument `bson:"programs"`
	}

	programsDocument struct {
		ID    string `bson:"id,omitempty"`
		Code  string `bson:"code"`
		Title string `bson:"title"`
	}
)

func newLevel(document levelDocument) *education.Level {
	level, _ := education.NewLevel(education.LevelParam{
		ID:    document.ID,
		Title: document.Title,
	})

	addUgsn(level, document.Ugsn)

	return level
}

func addUgsn(level *education.Level, document []ugsnDocument) {
	for _, v := range document {
		_ = level.AddUgsn(education.UgsnParams{
			ID:    v.ID,
			Code:  v.Code,
			Title: v.Title,
		})

		addSpecialty(level, v.Code, v.Specialties)
	}
}

func addSpecialty(level *education.Level, ucode string, documents []specialtiesDocument) {
	for _, v := range documents {
		_ = level.AddSpecialty(ucode, education.SpecialityParams{
			ID:       v.ID,
			Code:     v.Code,
			Title:    v.Title,
			UgsnCode: ucode,
		})

		addProgram(level, ucode, v.Code, v.Programs)
	}
}

func addProgram(level *education.Level, ucode string, scode string, documents []programsDocument) {
	for _, v := range documents {
		_ = level.AddProgram(ucode, scode, education.ProgramParams{
			ID:            v.ID,
			Code:          v.Code,
			Title:         v.Title,
			SpecialtyCode: scode,
		})
	}
}

func newLevelDocument(level *education.Level) levelDocument {
	return levelDocument{
		ID:    level.ID(),
		Title: level.Title(),
		Ugsn:  newUgsnDocument(level.AllUgsn()),
	}
}

func newUgsnDocument(ugsn []*education.Ugsn) []ugsnDocument {
	udocument := make([]ugsnDocument, 0, len(ugsn))

	for _, v := range ugsn {
		tmp := v

		udocument = append(udocument, ugsnDocument{
			ID:          tmp.ID(),
			Code:        tmp.Code(),
			Title:       tmp.Title(),
			Specialties: newSpecialtiesDocument(tmp.Specialities()),
		})
	}

	return udocument
}

func newSpecialtiesDocument(specialities []*education.Speciality) []specialtiesDocument {
	sdocument := make([]specialtiesDocument, 0, len(specialities))

	for _, v := range specialities {
		tmp := v

		sdocument = append(sdocument, specialtiesDocument{
			ID:       tmp.ID(),
			Code:     tmp.Code(),
			Title:    tmp.Title(),
			Programs: newProgramDocument(tmp.Programs()),
		})
	}

	return sdocument
}

func newProgramDocument(programs []*education.Program) []programsDocument {
	pdocument := make([]programsDocument, 0, len(programs))

	for _, v := range programs {
		tmp := v

		pdocument = append(pdocument, programsDocument{
			ID:    tmp.ID(),
			Code:  tmp.Code(),
			Title: tmp.Title(),
		})
	}

	return pdocument
}

func newLevelModelView(documents []levelDocument) []query.LevelModel {
	result := make([]query.LevelModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.LevelModel{
			ID:    document.ID,
			Title: document.Title,
		})
	}

	return result
}

func newUgsnModelView(documents []ugsnDocument) []query.UgsnModel {
	result := make([]query.UgsnModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.UgsnModel{
			ID:    document.ID,
			Code:  document.Code,
			Title: document.Title,
		})
	}

	return result
}

func newProgramModelViewWithDocuments(programs []programsDocument) []query.ProgramModel {
	result := make([]query.ProgramModel, 0, len(programs))

	for _, program := range programs {
		result = append(result, query.ProgramModel{
			ID:    program.ID,
			Code:  program.Code,
			Title: program.Title,
		})
	}

	return result
}

func newSpecialtiesModelViewWithDocuments(documents []specialtiesDocument) []query.SpecialtyModel {
	result := make([]query.SpecialtyModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.SpecialtyModel{
			ID:    document.ID,
			Code:  document.Code,
			Title: document.Title,
		})
	}

	return result
}

func newSpecificLevelView(document levelDocument) query.SpecificLevelModel {
	return query.SpecificLevelModel{
		ID:    document.ID,
		Title: document.Title,
		Ugsn:  newSpecificUgsnView(document.Ugsn),
	}
}

func newSpecificUgsnView(documents []ugsnDocument) []query.SpecificUgsnModel {
	result := make([]query.SpecificUgsnModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.SpecificUgsnModel{
			ID:          document.ID,
			Code:        document.Code,
			Title:       document.Title,
			Specialties: newSpecialtiesView(document.Specialties),
		})
	}

	return result
}

func newSpecialtiesView(documents []specialtiesDocument) []query.SpecificSpecialtyModel {
	result := make([]query.SpecificSpecialtyModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.SpecificSpecialtyModel{
			ID:       document.ID,
			Code:     document.Code,
			Title:    document.Title,
			Programs: newProgramView(document.Programs),
		})
	}

	return result
}

func newProgramView(documents []programsDocument) []query.ProgramModel {
	result := make([]query.ProgramModel, 0, len(documents))

	for _, document := range documents {
		result = append(result, query.ProgramModel{
			ID:    document.ID,
			Code:  document.Code,
			Title: document.Title,
		})
	}

	return result
}
