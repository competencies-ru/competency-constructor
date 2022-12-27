// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package v1

// Defines values for ErrorSlug.
const (
	BadRequest                   ErrorSlug = "bad-request"
	EmptyBearerToken             ErrorSlug = "empty-bearer-token"
	InvalidJson                  ErrorSlug = "invalid-json"
	InvalidLevelParameters       ErrorSlug = "invalid-level-parameters"
	InvalidProgramParameters     ErrorSlug = "invalid-program-parameters"
	InvalidSpecialtiesParameters ErrorSlug = "invalid-specialties-parameters"
	InvalidUgsnParameters        ErrorSlug = "invalid-ugsn-parameters"
	LevelNotFound                ErrorSlug = "level-not-found"
	ProgramNotFound              ErrorSlug = "program-not-found"
	SpecialtyNotFound            ErrorSlug = "specialty-not-found"
	UgsnNotFound                 ErrorSlug = "ugsn-not-found"
	UnableToVerifyJwt            ErrorSlug = "unable-to-verify-jwt"
	UnauthorizedUser             ErrorSlug = "unauthorized-user"
	UnexpectedError              ErrorSlug = "unexpected-error"
)

// CreateLevelRequest defines model for CreateLevelRequest.
type CreateLevelRequest struct {
	Title string `json:"title"`
}

// CreateProgramRequest defines model for CreateProgramRequest.
type CreateProgramRequest struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// CreateSpecialtyRequest defines model for CreateSpecialtyRequest.
type CreateSpecialtyRequest struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// CreateUgsnRequest defines model for CreateUgsnRequest.
type CreateUgsnRequest struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// Error defines model for Error.
type Error struct {
	Details string    `json:"details"`
	Slug    ErrorSlug `json:"slug"`
}

// ErrorSlug defines model for ErrorSlug.
type ErrorSlug string

// LevelResponse defines model for LevelResponse.
type LevelResponse struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// ProgramResponse defines model for ProgramResponse.
type ProgramResponse struct {
	Code  string `json:"code"`
	Id    string `json:"id"`
	Title string `json:"title"`
}

// SpecialtyResponse defines model for SpecialtyResponse.
type SpecialtyResponse struct {
	Code  string `json:"code"`
	Id    string `json:"id"`
	Title string `json:"title"`
}

// SpecificLevelResponse defines model for SpecificLevelResponse.
type SpecificLevelResponse struct {
	Id    string                 `json:"id"`
	Title string                 `json:"title"`
	Ugsn  []SpecificUgsnResponse `json:"ugsn"`
}

// SpecificSpecialtyResponse defines model for SpecificSpecialtyResponse.
type SpecificSpecialtyResponse struct {
	Code    string            `json:"code"`
	Id      string            `json:"id"`
	Program []ProgramResponse `json:"program"`
	Title   string            `json:"title"`
}

// SpecificUgsnResponse defines model for SpecificUgsnResponse.
type SpecificUgsnResponse struct {
	Code      string                      `json:"code"`
	Id        string                      `json:"id"`
	Specialty []SpecificSpecialtyResponse `json:"specialty"`
	Title     string                      `json:"title"`
}

// UgsnResponse defines model for UgsnResponse.
type UgsnResponse struct {
	Code  string `json:"code"`
	Id    string `json:"id"`
	Title string `json:"title"`
}

// AddUgsnJSONBody defines parameters for AddUgsn.
type AddUgsnJSONBody = []CreateUgsnRequest

// AddSpecialtiesJSONBody defines parameters for AddSpecialties.
type AddSpecialtiesJSONBody = []CreateSpecialtyRequest

// AddProgramsJSONBody defines parameters for AddPrograms.
type AddProgramsJSONBody = []CreateProgramRequest

// CreateLevelJSONRequestBody defines body for CreateLevel for application/json ContentType.
type CreateLevelJSONRequestBody = CreateLevelRequest

// AddUgsnJSONRequestBody defines body for AddUgsn for application/json ContentType.
type AddUgsnJSONRequestBody = AddUgsnJSONBody

// AddSpecialtiesJSONRequestBody defines body for AddSpecialties for application/json ContentType.
type AddSpecialtiesJSONRequestBody = AddSpecialtiesJSONBody

// AddProgramsJSONRequestBody defines body for AddPrograms for application/json ContentType.
type AddProgramsJSONRequestBody = AddProgramsJSONBody
