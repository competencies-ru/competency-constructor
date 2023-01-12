package v1

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/competencies-ru/competency-constructor/internal/core/app/query"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"github.com/competencies-ru/competency-constructor/internal/core/adapter/driver/rest"
	"github.com/competencies-ru/competency-constructor/internal/core/app/command"
	"github.com/competencies-ru/competency-constructor/internal/core/entity/competencies"
)

func decodeCreateCompetencyRequest(w http.ResponseWriter, r *http.Request) (command.CreateCompetencyCommand, bool) {
	var request CreateCompetencyRequest
	if !decode(w, r, &request) {
		return command.CreateCompetencyCommand{}, false
	}

	ctype, ok := decodeCompetencyType(w, r, request.Type)

	if !ok {
		return command.CreateCompetencyCommand{}, ok
	}

	return command.CreateCompetencyCommand{
		Code:           request.Code,
		Title:          request.Title,
		Category:       request.Category,
		CompetencyType: ctype,
		LevelID:        toString(request.LevelId),
		UgsnID:         toString(request.UgsnId),
		SpecialtyID:    toString(request.SpecialtyId),
		ProgramID:      toString(request.ProgramId),
	}, true
}

func decodeCompetencyType(w http.ResponseWriter, r *http.Request, ctype CompetencyType) (competencies.Type, bool) {
	switch ctype {
	case UNIVERSAL:
		return competencies.UNIVERSAL, true
	case GENERAL:
		return competencies.GENERAL, true
	case PROFESSIONAL:
		return competencies.PROFESSIONAL, true
	}

	rest.UnprocessableEntity(string(InvalidCompetenciesParameters), competencies.ErrNoSuchValue, w, r)

	return competencies.Type(0), false
}

func mappingFilterCompetencyParams(
	w http.ResponseWriter,
	r *http.Request,
	param FilterCompetencyParams,
) (query.FilterCompetencyParam, bool) {
	if !validateParam(param) {
		rest.BadRequest(string(BadRequest), errors.New("all params is nil"), w, r)

		return query.FilterCompetencyParam{}, false
	}

	return query.FilterCompetencyParam{
		LevelID:     toString(param.LevelId),
		UgsnID:      toString(param.UgsnId),
		ProgramID:   toString(param.ProgramId),
		SpecialtyID: toString(param.SpecialtyId),
	}, true
}

func validateParam(param FilterCompetencyParams) bool {
	return param.LevelId != nil ||
		param.UgsnId != nil ||
		param.SpecialtyId != nil ||
		param.ProgramId != nil
}

func mappingCompetencyType(t competencies.Type) CompetencyType {
	switch t {
	case competencies.UNIVERSAL:
		return UNIVERSAL
	case competencies.PROFESSIONAL:
		return PROFESSIONAL
	case competencies.GENERAL:
		return GENERAL
	}

	return ""
}

func renderCompetenciesResponse(w http.ResponseWriter, r *http.Request, models []query.CompetencyModel) {
	response := make([]CompetencyResponse, 0, len(models))

	for _, model := range models {
		response = append(response, CompetencyResponse{
			Category:    model.Category,
			Code:        model.Code,
			Id:          uuid.MustParse(model.ID),
			LevelId:     toUUIDPointer(model.LevelID),
			ProgramId:   toUUIDPointer(model.ProgramID),
			SpecialtyId: toUUIDPointer(model.SpecialtyID),
			Title:       model.Title,
			Type:        mappingCompetencyType(model.CompetencyType),
			UgsnId:      toUUIDPointer(model.UgsnID),
		})
	}

	render.Respond(w, r, response)
}
