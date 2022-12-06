package service

type (
	SpecificUgsn struct {
		Code              string
		Title             string
		SpecificSpecialty []SpecificSpecialty
	}

	SpecificSpecialty struct {
		Code     string
		Title    string
		UgsnCode string
		Programs []Program
	}

	Program struct {
		ID            string
		Title         string
		SpecialtyCode string
	}

	Specialty struct {
		Code  string
		Title string
	}

	Specialties []Specialty
)
