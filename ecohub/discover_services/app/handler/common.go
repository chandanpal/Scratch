package handler

import (
	"encoding/json"
	"net/http"
	_"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	_"io"
	_"github.com/discover_services/app/model"
)

// respondJSON makes the response with payload as json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// respondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
	
}


func NotImplemented(db *mongo.Client, w http.ResponseWriter, r *http.Request) {

	respondJSON(w, http.StatusOK, "API Not Implemented")
}


func getResponseData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	return contents, nil
}
