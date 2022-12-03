package v1

import "net/http"

func (h handler) GetSpecificSpecialty(w http.ResponseWriter, r *http.Request, specialtyCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) GetPrograms(w http.ResponseWriter, r *http.Request, specialtyCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) AddedProgram(w http.ResponseWriter, r *http.Request, specialtyCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) GetUgsn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) CreateUgsn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) GetSpecificUgsn(w http.ResponseWriter, r *http.Request, ugsnCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) GetSpecialties(w http.ResponseWriter, r *http.Request, ugsnCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h handler) AddedSpecialties(w http.ResponseWriter, r *http.Request, ugsnCode string) {
	w.WriteHeader(http.StatusNotImplemented)
}
