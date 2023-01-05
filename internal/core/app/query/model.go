package query

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
