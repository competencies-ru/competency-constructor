package query

import "github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"

type (
	LevelModel struct {
		ID    string
		Title string
	}

	UgsnModel struct {
		ID      string
		Code    string
		Title   string
		LevelID string
	}

	SpecialtyModel struct {
		ID     string
		Code   string
		Title  string
		UgsnID string
	}

	SpecificLevelModel struct {
		ID    string
		Title string
		Ugsn  []SpecificUgsnModel
	}
)

type (
	SpecificSpecialtyModel struct {
		ID       string
		Code     string
		Title    string
		Programs []ProgramModel
	}

	ProgramModel struct {
		ID          string
		Code        string
		Title       string
		SpecialtyID string
	}
)

type (
	SpecificUgsnModel struct {
		ID          string
		Code        string
		Title       string
		Specialties []SpecificSpecialtyModel
	}
)

type (
	FilterCompetencyParam struct {
		LevelID     string
		UgsnID      string
		ProgramID   string
		SpecialtyID string
	}

	CompetencyModel struct {
		ID             string
		Code           string
		Title          string
		Category       string
		CompetencyType competencies.Type
		LevelID        string
		UgsnID         string
		SpecialtyID    string
		ProgramID      string
	}
)
