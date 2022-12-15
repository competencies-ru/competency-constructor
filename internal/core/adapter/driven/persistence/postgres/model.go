package postgres

import (
	"database/sql"
	"github.com/competencies-ru/competency-constructor/internal/core/app/service"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/specialty"
	"github.com/competencies-ru/competency-constructor/internal/utils"
	"github.com/google/uuid"
	"go.uber.org/multierr"
)

type (
	ugsn struct {
		Code  string
		Title string
	}

	specificUgsn struct {
		Code        string
		Title       string
		Specialties map[string]*specificSpecialty
	}

	specificSpecialty struct {
		Code     string
		Title    string
		UgsnCode string
		Programs map[string]*program
	}

	program struct {
		ID            string
		Title         string
		SpecialtyCode string
	}
)

func mapUgsn(input []*ugsn) ([]*specialty.Ugsn, error) {
	var result error

	output := make([]*specialty.Ugsn, 0, len(input))

	for i := range input {
		v := input[i]

		newUgsn, err := specialty.NewUgsn(specialty.UgsnParams{
			Code:  v.Code,
			Title: v.Title,
		})
		if err != nil {
			result = multierr.Append(result, err)

			continue
		}

		output = append(output, newUgsn)
	}

	return output, result
}

func unmarshallingSpecificUgsn(su *specificUgsn, scode, stitle, sugsn, pid, ptitle, pspecialty sql.NullString) {
	if !scode.Valid {
		return
	}

	if _, ok := su.Specialties[scode.String]; !ok {
		su.Specialties[scode.String] = &specificSpecialty{
			Code:     scode.String,
			Title:    stitle.String,
			UgsnCode: sugsn.String,
			Programs: make(map[string]*program),
		}
	}

	s := su.Specialties[scode.String]

	unmarshallingSpecificProgram(s, pid, ptitle, pspecialty)
}

func unmarshallingSpecificProgram(s *specificSpecialty, pid, ptitle, pspecialty sql.NullString) {
	if !pid.Valid {
		return
	}

	if _, ok := s.Programs[pid.String]; !ok {
		s.Programs[pid.String] = &program{
			ID:            pid.String,
			Title:         ptitle.String,
			SpecialtyCode: pspecialty.String,
		}
	}
}

func marshallingSpecificUgsn(s *specificUgsn) *service.SpecificUgsn {
	return &service.SpecificUgsn{
		Code:              s.Code,
		Title:             s.Title,
		SpecificSpecialty: utils.ConvertMapPointer(s.Specialties, mapS),
	}
}

func mapS(s *specificSpecialty) *service.SpecificSpecialty {
	mapper := func(t *program) *service.Program {

		id, _ := uuid.Parse(t.ID)

		return &service.Program{
			ID:            id,
			Title:         t.Title,
			SpecialtyCode: t.SpecialtyCode,
		}
	}

	return &service.SpecificSpecialty{
		Code:     s.Code,
		Title:    s.Title,
		UgsnCode: s.UgsnCode,
		Programs: utils.ConvertMapPointer(s.Programs, mapper),
	}
}
