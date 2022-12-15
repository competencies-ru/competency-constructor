package v1

import "net/http"

func (h handler) GetPrograms(w http.ResponseWriter, r *http.Request, specialtyCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) AddedProgram(w http.ResponseWriter, r *http.Request, specialtyCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) GetSpecificSpecialty(w http.ResponseWriter, r *http.Request, specialtyCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}
