package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func writeJson(w http.ResponseWriter, container any) {
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(container)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(j)
}

func readJson(w http.ResponseWriter, r *http.Request, container any) error {
	err := json.NewDecoder(r.Body).Decode(container)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

func getQueryParamInt(param string, defaultValue int, r *http.Request) int {
	p := r.URL.Query().Get(param)
	res, err := strconv.Atoi(p)
	if err != nil {
		return defaultValue
	}
	return res
}

func getQueryParamString(param string, r *http.Request) string {
	p := r.URL.Query().Get(param)
	return p
}
