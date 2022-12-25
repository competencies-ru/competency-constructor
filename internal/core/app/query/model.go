package query

type (
	LevelModel struct {
		ID    string
		Title string
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
		ID    string
		Code  string
		Title string
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
