package v1

import (
	"net/http"

	"github.com/competencies-ru/competency-constructor/internal/core/app/service"

	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/go-chi/render"
)

func renderUgsnResponse(
	w http.ResponseWriter,
	r *http.Request,
	ugsn []*specialty.Ugsn,
) {
	response := make([]UgsnResponse, 0, len(ugsn))
	for _, v := range ugsn {
		response = append(response, marshalUgsnTo(*v))
	}

	render.Respond(w, r, response)
}

func marshalUgsnTo(ugsn specialty.Ugsn) UgsnResponse {
	return UgsnResponse{
		Code:  ugsn.Code(),
		Title: ugsn.Title(),
	}
}

func unmarshallingCreateUgsn(w http.ResponseWriter, r *http.Request) (uc service.UgsnCreate, ok bool) {
	if ok = decode(w, r, &uc); !ok {
		return
	}

	return service.UgsnCreate{
		Code:  uc.Code,
		Title: uc.Title,
	}, true
}

func renderSpecificUgsn(w http.ResponseWriter, r *http.Request, ugsn *service.SpecificUgsn) {

	specialties := unmarshallingSpecialties(ugsn.SpecificSpecialty)

	response := SpecificUgsnResponse{
		Code:      ugsn.Code,
		Specialty: specialties,
		Title:     ugsn.Title,
	}

	render.Respond(w, r, response)
}

func unmarshallingSpecialties(s []*service.SpecificSpecialty) []SpecificSpecialtyResponse {
	sp := make([]SpecificSpecialtyResponse, 0, len(s))

	for _, v := range s {
		sp = append(sp, unmarshallingSpecialty(v))
	}

	return sp
}

func unmarshallingSpecialty(s *service.SpecificSpecialty) SpecificSpecialtyResponse {
	programs := unmarshallingPrograms(s.Programs)

	return SpecificSpecialtyResponse{
		Code:     s.Code,
		UgsnCode: s.UgsnCode,
		Title:    s.Title,
		Program:  programs,
	}
}

func unmarshallingPrograms(p []*service.Program) []ProgramResponse {

	pr := make([]ProgramResponse, 0, len(p))

	for _, v := range p {
		pr = append(pr, ProgramResponse{
			Id:            v.ID,
			SpecialtyCode: v.SpecialtyCode,
			Title:         v.Title,
		})
	}

	return pr
}
