package controllers

import (
	"encoding/json"
	"net/http"

	dao "../dao"
)

// TestController : for test purpose
func TestController(w http.ResponseWriter, r *http.Request) {
	prj := dao.QueryProject(39)
	respondWithJSON(w, http.StatusOK, prj)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
