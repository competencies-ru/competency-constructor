// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.3 DO NOT EDIT.
package v1

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns level of the educational program.
	// (GET /levels)
	GetLevels(w http.ResponseWriter, r *http.Request)
	// Create the level of the educational program
	// (POST /levels)
	CreateLevel(w http.ResponseWriter, r *http.Request)
	// Returns level with such ID.
	// (GET /levels/{levelId})
	GetSpecificLevel(w http.ResponseWriter, r *http.Request, levelId string)
	// return ugsn by level id
	// (GET /levels/{levelId}/ugsn)
	GetUgsn(w http.ResponseWriter, r *http.Request, levelId string)
	// addedUgsn by level id
	// (POST /levels/{levelId}/ugsn)
	AddUgsn(w http.ResponseWriter, r *http.Request, levelId string)
	// delete program by id
	// (DELETE /programs/{id})
	DeleteProgram(w http.ResponseWriter, r *http.Request, id string)
	// delete specialty by code ugsn and level id
	// (DELETE /specialty/{id})
	DeleteSpecialty(w http.ResponseWriter, r *http.Request, id string)
	// return specific ugsn by level id
	// (GET /specialty/{id})
	GetSpecificSpecialty(w http.ResponseWriter, r *http.Request, id string)
	// return programs
	// (GET /specialty/{id}/programs)
	GetPrograms(w http.ResponseWriter, r *http.Request, id string)
	// create programs
	// (POST /specialty/{id}/programs)
	AddPrograms(w http.ResponseWriter, r *http.Request, id string)
	// delete ugsn by id
	// (DELETE /ugsn/{id})
	DeleteUgsn(w http.ResponseWriter, r *http.Request, id string)
	// return specific ugsn by level id
	// (GET /ugsn/{id})
	GetSpecificUgsn(w http.ResponseWriter, r *http.Request, id string)
	// return specialties by ugsn id
	// (GET /ugsn/{id}/specialties)
	GetSpecialties(w http.ResponseWriter, r *http.Request, id string)
	// add specialties by level id and ugsn code
	// (POST /ugsn/{id}/specialties)
	AddSpecialties(w http.ResponseWriter, r *http.Request, id string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetLevels operation middleware
func (siw *ServerInterfaceWrapper) GetLevels(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetLevels(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// CreateLevel operation middleware
func (siw *ServerInterfaceWrapper) CreateLevel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateLevel(w, r)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSpecificLevel operation middleware
func (siw *ServerInterfaceWrapper) GetSpecificLevel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "levelId" -------------
	var levelId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "levelId", runtime.ParamLocationPath, chi.URLParam(r, "levelId"), &levelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "levelId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSpecificLevel(w, r, levelId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUgsn operation middleware
func (siw *ServerInterfaceWrapper) GetUgsn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "levelId" -------------
	var levelId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "levelId", runtime.ParamLocationPath, chi.URLParam(r, "levelId"), &levelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "levelId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUgsn(w, r, levelId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddUgsn operation middleware
func (siw *ServerInterfaceWrapper) AddUgsn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "levelId" -------------
	var levelId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "levelId", runtime.ParamLocationPath, chi.URLParam(r, "levelId"), &levelId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "levelId", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddUgsn(w, r, levelId)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteProgram operation middleware
func (siw *ServerInterfaceWrapper) DeleteProgram(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteProgram(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteSpecialty operation middleware
func (siw *ServerInterfaceWrapper) DeleteSpecialty(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteSpecialty(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSpecificSpecialty operation middleware
func (siw *ServerInterfaceWrapper) GetSpecificSpecialty(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSpecificSpecialty(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetPrograms operation middleware
func (siw *ServerInterfaceWrapper) GetPrograms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPrograms(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddPrograms operation middleware
func (siw *ServerInterfaceWrapper) AddPrograms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddPrograms(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteUgsn operation middleware
func (siw *ServerInterfaceWrapper) DeleteUgsn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUgsn(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSpecificUgsn operation middleware
func (siw *ServerInterfaceWrapper) GetSpecificUgsn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSpecificUgsn(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetSpecialties operation middleware
func (siw *ServerInterfaceWrapper) GetSpecialties(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetSpecialties(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// AddSpecialties operation middleware
func (siw *ServerInterfaceWrapper) AddSpecialties(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	var handler http.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AddSpecialties(w, r, id)
	})

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshallingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshallingParamError) Error() string {
	return fmt.Sprintf("Error unmarshalling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshallingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/levels", wrapper.GetLevels)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/levels", wrapper.CreateLevel)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/levels/{levelId}", wrapper.GetSpecificLevel)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/levels/{levelId}/ugsn", wrapper.GetUgsn)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/levels/{levelId}/ugsn", wrapper.AddUgsn)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/programs/{id}", wrapper.DeleteProgram)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/specialty/{id}", wrapper.DeleteSpecialty)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/specialty/{id}", wrapper.GetSpecificSpecialty)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/specialty/{id}/programs", wrapper.GetPrograms)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/specialty/{id}/programs", wrapper.AddPrograms)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/ugsn/{id}", wrapper.DeleteUgsn)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/ugsn/{id}", wrapper.GetSpecificUgsn)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/ugsn/{id}/specialties", wrapper.GetSpecialties)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/ugsn/{id}/specialties", wrapper.AddSpecialties)
	})

	return r
}
